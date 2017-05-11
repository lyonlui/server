package gate

import (
	"server/login"
	"server/msg"
)

func init() {

	msg.Processor.SetRouter(&msg.Login{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.AchieveSession{}, login.ChanRPC)

}
