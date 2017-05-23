package gate

import (
	"server/login"
	"server/msg"
)

func init() {

	msg.Processor.SetRouter(&msg.Login{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.LoginFeedback{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.Quit{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.QuitFeedback{}, login.ChanRPC)

}
