package internal

import (
	"fmt"
	"server/users"

	"github.com/name5566/leaf/gate"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a
	fmt.Println("rpcNewAgent")
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)

	if a.UserData().Verify {
		users.DeleteAgent(a.UserData().UserID)
	}
	fmt.Println("all users: ", users.GetAgentCounts())
	fmt.Println("rpcCloseAgent")
}
