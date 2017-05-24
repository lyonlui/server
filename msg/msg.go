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

type Login struct {
	Accounts      string
	Password      string
	ClientSerial  string
	Device        string
	ClientVersion string
}

type LoginFeedback struct {
	LoginStatus   bool
	ErrorDescribe string
	NickName      string
	UnderWrite    string
	FaceID        int
	Gender        int
	CustomFaceVer int
}

type Quit struct {
}

type QuitFeedback struct {
	QuitStatus bool
}
type Warnning struct {
}
