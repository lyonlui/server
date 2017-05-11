package msg

import (
	"github.com/name5566/leaf/network/json"
)

var Processor = json.NewProcessor()

func init() {
	Processor.Register(&Login{})
	Processor.Register(&AchieveSession{})
	Processor.Register(&LoginFeedback{})
	Processor.Register(&Session{})
}

type Login struct {
	Accounts      string
	Password      string
	ClientSerial  string
	Device        int
	ClientVersion string
}

type AchieveSession struct {
}

type LoginFeedback struct {
	LoginStatus bool
	ErrorCode   int
}

type Session struct {
	Status    bool
	SessionID string
}
