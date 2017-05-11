package internal

import (
	"reflect"
	"server/msg"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&msg.Login{}, handleLogin)
	handleMsg(&msg.Session{}, handleSession)

}

func handleLogin(args []interface{}) {

	// 收到的 Login 消息
	m := args[0].(*msg.Login)
	// 消息的发送者
	a := args[1].(gate.Agent)

	// 输出收到的消息的内容
	log.Debug("hello %v", m.Accounts)

	//判断用户名是否存在

	//判断密码是否正确

	//判断用户是否已在线

	// 给发送者回应一个 Hello 消息
	a.WriteMsg(&msg.LoginFeedback{
		LoginStatus: true,
		ErrorCode:   100,
	})

}

func handleSession(args []interface{}) {

}
