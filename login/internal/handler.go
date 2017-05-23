package internal

import (
	"fmt"
	"reflect"
	"server/msg"
	"strings"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

const TAG string = "server/login/internal/handler.go"

var db *sql.DB

type AccountInfo struct {
	UserID        int
	Accounts      string
	NickName      string
	UnderWrite    string
	FaceID        int
	Gender        int
	CustomFaceVer int
}

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {

	handleMsg(&msg.Login{}, handleLogin)
	handleMsg(&msg.Quit{}, handleQuit)

}

func handleLogin(args []interface{}) {

	var loginStatus bool

	// 收到的 Login 消息
	userMag := args[0].(*msg.Login)
	// 消息的发送者
	userAgent := args[1].(gate.Agent)

	// 输出收到的消息的内容
	log.Debug("hello %v", userMag.Accounts)

	//消息字段过滤，防注入
	loginStatus, strErrorDescribe, accInfo := loginVerify(userMag.Accounts, userMag.Password, strings.Split(userAgent.RemoteAddr().String(), ":")[0], userMag.ClientSerial)

	fmt.Println(strErrorDescribe)

	//设置用户session

	if loginStatus {
		userData := userAgent.UserData()
		userData.UserID = accInfo.UserID
		userData.Verify = true
		userAgent.SetUserData(userData)
	}

	// 给发送者回应一个 Hello 消息
	userAgent.WriteMsg(&msg.LoginFeedback{
		LoginStatus:   loginStatus,
		ErrorDescribe: strErrorDescribe,
		NickName:      accInfo.NickName,
		UnderWrite:    accInfo.UnderWrite,
		FaceID:        accInfo.FaceID,
		Gender:        accInfo.Gender,
		CustomFaceVer: accInfo.CustomFaceVer,
	})

}

// 收到的 Quit 消息
func handleQuit(args []interface{}) {

	// 消息的发送者
	userAgent := args[1].(gate.Agent)

	userAgent.WriteMsg(&msg.QuitFeedback{
		QuitStatus: false,
	})

	userAgent.Close()

}

/****************************************************************************

	登录验证函数


*****************************************************************************/

func loginVerify(strAccounts string, strPassword string, strClientIP string, machineSerial string) (loginStatus bool, strErrorDescribe string, accInfo AccountInfo) {

	var info AccountInfo

	db, e := sql.Open("mysql", "root:HelloWorld!@/AccountDB")
	if e != nil {
		fmt.Println("connect db error")
		return false, "", info
	}
	defer db.Close()

	//系统暂停
	var StatusValue int
	row := db.QueryRow("SELECT StatusValue FROM SystemStatusInfo WHERE StatusName = ?", "EnjoinLogon")
	err := row.Scan(&StatusValue)
	if err != sql.ErrNoRows {
		if StatusValue != 0 {
			row := db.QueryRow("SELECT StatusDescription FROM SystemStatusInfo WHERE StatusName = ?", "EnjoinLogon")
			err := row.Scan(&strErrorDescribe)
			if err == sql.ErrNoRows {
				strErrorDescribe = "系统维护中"
			}
			return false, strErrorDescribe, info
		}
	}

	//效验地址
	var EnjoinLogin int
	row = db.QueryRow("SELECT EnjoinLogin FROM ConfineAddress WHERE AddrString = ? AND (EnjoinOverDate>NOW() OR EnjoinOverDate IS NULL)", strClientIP)
	err = row.Scan(&EnjoinLogin)
	fmt.Println(err)
	if err != sql.ErrNoRows {
		if EnjoinLogin == 1 {
			strErrorDescribe = "抱歉地通知您，系统禁止了您所在的 IP 地址的登录功能，请联系客户服务中心了解详细情况！"
			fmt.Println(strErrorDescribe)
			return false, strErrorDescribe, info
		}
	}

	//查询用户
	var UserID, Nullity, StunDown, FaceID, Gender, CustomFaceVer, PlayTimeCount, MoorMachine int
	var NickName, UnderWrite, LoginPass, Accounts, MachineSerial string

	stmt, _ := db.Prepare(`SELECT UserID,Accounts,NickName,UnderWrite,LoginPass,FaceID,Gender,Nullity,StunDown,CustomFaceVer,PlayTimeCount,MoorMachine,MachineSerial FROM AccountsInfo WHERE Accounts = ?`)
	row = stmt.QueryRow(strAccounts)
	defer stmt.Close()
	err = row.Scan(&UserID, &Accounts, &NickName, &UnderWrite, &LoginPass, &FaceID, &Gender, &Nullity, &StunDown, &CustomFaceVer, &PlayTimeCount, &MoorMachine, &MachineSerial)

	if err == sql.ErrNoRows {

		strErrorDescribe = "您的帐号不存在或者密码输入有误，请查证后再次尝试登录！"
		return false, strErrorDescribe, info
	}

	if Nullity == 1 {
		strErrorDescribe = "您的帐号暂时处于冻结状态，请联系客户服务中心了解详细情况！"
		return false, strErrorDescribe, info
	}

	if StunDown == 1 {
		strErrorDescribe = "您的帐号使用了安全关闭功能，必须重新开通后才能继续使用！"
		return false, strErrorDescribe, info
	}

	if MoorMachine == 1 {
		if MachineSerial != machineSerial {
			strErrorDescribe = "您的帐号使用固定机器登陆功能，您现所使用的机器不是所指定的机器！"
			return false, strErrorDescribe, info
		}
	}

	if LoginPass != strPassword {
		strErrorDescribe = "您的帐号不存在或者密码输入有误，请查证后再次尝试登录！"
		return false, strErrorDescribe, info
	} else {
		strErrorDescribe = "登录成功"
	}

	if MoorMachine == 2 {
		db.Exec("UPDATE AccountsInfo SET MoorMachine = 1 ,MachineSerial = ? WHERE = UserID=?", machineSerial, UserID)
		strErrorDescribe = "您的帐号成功使用了固定机器登陆功能！"
	}

	//更新信息
	stmt, _ = db.Prepare(`UPDATE AccountsInfo SET GameLoginTimes = GameLoginTimes+1 , LastLoginDate =NOW() , LastLoginIP = ? ,LastLoginMachine = ? WHERE UserID=?`)
	stmt.Exec(strClientIP, machineSerial, UserID)
	fmt.Println("here")

	//记录日志
	db.Exec("INSERT INTO SystemStreamInfo (DateID, GameLoginSuccess) VALUES (unix_timestamp(curdate()),1)  ON DUPLICATE KEY UPDATE GameLoginSuccess=GameLoginSuccess+1")

	fmt.Println("here")
	fmt.Println(strAccounts, strPassword, strClientIP)

	//输出变量

	info.Accounts = Accounts
	info.NickName = NickName
	info.CustomFaceVer = CustomFaceVer
	info.FaceID = FaceID
	info.Gender = Gender
	info.UnderWrite = UnderWrite
	info.UserID = UserID

	return true, strErrorDescribe, info
}
