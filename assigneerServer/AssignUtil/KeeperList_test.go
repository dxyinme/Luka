package AssignUtil

import (
	"github.com/dxyinme/LukaComm/chatMsg"
	"github.com/dxyinme/LukaComm/util"
	"log"
	"testing"
)

func TestKeeperList_ToBytes(t *testing.T) {
	k := &KeeperList{
		KeeperId: make([]uint32,1),
		Lis:      make([]*chatMsg.KeepAlive,1),
	}
	k.Lis[0] = &chatMsg.KeepAlive{
		CheckAlive:  "hehe",
		Errors:      make([]string,1),
		MsgRecv:     1,
		MsgSend:     10,
		MsgNotLocal: 100,
	}
	k.KeeperId[0] = 1

	b := k.ToBytes()
	log.Println(string(b))
	var jsb = &KeeperList{}
	err := util.IJson.Unmarshal(b,jsb)
	if err != nil {
		t.Error(err)
	} else {
		log.Println(jsb)
	}
}