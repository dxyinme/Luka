package WorkerPool

import (
	"github.com/dxyinme/LukaComm/chatMsg"
	"log"
	"testing"
)

func TestCopy(t *testing.T) {
	T1 := &chatMsg.Msg{
		From:           "",
		Target:         "",
		Content:        nil,
		MsgType:        0,
		MsgContentType: 0,
		SendTime:       "",
		GroupName:      "",
		Spread:         false,
	}
	T2 := *T1
	log.Printf("%p\n", T1)
	log.Printf("%p\n", &T2)
	T2.Spread = true
	if T1.Spread {
		t.Fatal("error")
	}
}