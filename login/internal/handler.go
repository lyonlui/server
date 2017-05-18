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
	Nickname      string
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
	handleMsg(&msg.Session{}, handleSession)

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

	// 给发送者回应一个 Hello 消息
	userAgent.WriteMsg(&msg.LoginFeedback{
		LoginStatus:   loginStatus,
		ErrorDescribe: strErrorDescribe,
		AccountInfo:   accInfo,
	})

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
			fmt.Println(err)
			fmt.Println(strErrorDescribe)
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
	var UserID, Nullity, StunDown, FaceID, Gender, CustomFaceVer, PlayTimeCount int
	var NickName, UnderWrite, LoginPass, Accounts string
	row = db.QueryRow("SELECT UserID,Accounts,NickName,UnderWrite,LoginPass,FaceID,Gender,Nullity,StunDown,CustomFaceVer,PlayTimeCount FROM AccountsInfo WHERE Accounts = ?", strAccounts)
	err = row.Scan(&UserID, &Accounts, &NickName, &UnderWrite, &LoginPass, &FaceID, &Gender, &Nullity, &StunDown, &CustomFaceVer, &PlayTimeCount)

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

	if LoginPass != strPassword {
		strErrorDescribe = "您的帐号不存在或者密码输入有误，请查证后再次尝试登录！"
		return false, strErrorDescribe, info
	}

	//更新信息
	db.Exec("UPDATE AccountsInfo SET GameLoginTimes = GameLoginTimes+1 , `LastLoginDate`=NOW() , `LastLoginIP` = ? ,`LastLoginMachine` = ? WHERE UserID=?", strClientIP, machineSerial, UserID)

	fmt.Println("here")

	//记录日志
	db.Exec("INSERT INTO SystemStreamInfo (DateID, GameLoginSuccess) VALUES (unix_timestamp(curdate()),1)  ON DUPLICATE KEY UPDATE GameLoginSuccess=GameLoginSuccess+1")

	fmt.Println("here")
	fmt.Println(strAccounts, strPassword, strClientIP)

	//输出变量
	strErrorDescribe = "登录成功"
	info.Accounts = Accounts
	info.Nickname = NickName
	info.CustomFaceVer = CustomFaceVer
	info.FaceID = FaceID
	info.Gender = Gender
	info.UnderWrite = UnderWrite
	info.UserID = UserID

	return true, strErrorDescribe, info
}

func handleSession(args []interface{}) {

}
