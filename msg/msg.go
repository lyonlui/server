package msg

import (
	"github.com/name5566/leaf/network/json"
)

var Processor = json.NewProcessor()

func init() {
	Processor.Register(&Login{})
	Processor.Register(&LoginFeedback{})
	Processor.Register(&Quit{})
	Processor.Register(&QuitFeedback{})
	Processor.Register(&Warnning{})

}

//服务器相应消息

/*****用户登录请求消息*********************************************/
type Login struct {
	Accounts      string
	Password      string
	ClientSerial  string
	Device        string
	ClientVersion string
}

/******用户登录请求返回消息*****************************************/
type LoginFeedback struct {
	LoginStatus   bool
	ErrorDescribe string
	NickName      string
	UnderWrite    string
	FaceID        int
	Gender        int
	CustomFaceVer int
}

/******用户获取头像消息**********************************************/
type AchieveLogo struct {
}

/******返回头像路径**************************************************/
type LogoPath struct {
	LogoPath string
}

/*******用户退出登录请求**********************************************/
type Quit struct {
}

/*******用户退出请求返回**********************************************/
type QuitFeedback struct {
	QuitStatus bool
}

/*******警告消息****************************************************/
type Warnning struct {
	WarnDescribe string
}

////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////
//////////////////////////游戏模块消息////////////////////////////////

//加入游戏
type JoinGame struct {
}

//加入游戏反馈
type JoinGameFeedback struct {
}

//当前游戏进行的状态

//是否之前有本局已下注信息

//下注消息
type Cathectic struct {
}

//下注是否成功消息
type CathecticFeedback struct {
}

//本局结果消息
type GameResult struct {
}

//下一局开始消息

//从新开始一轮消息
