package AssignUtil

import (
	"github.com/dxyinme/LukaComm/chatMsg"
	"github.com/dxyinme/LukaComm/util"
)

type KeeperList struct {
	KeeperId 	[]uint32
	Lis 		[]*chatMsg.KeepAlive
}

func (k *KeeperList) ToBytes() []byte {
	if k == nil {
		return nil
	} else {
		b,err := util.IJson.Marshal(k)
		if err != nil {
			return nil
		}
		return b
	}
}