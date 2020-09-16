package WorkerPool

import "github.com/dxyinme/LukaComm/chatMsg"

type NormalImpl struct {

}

func (ni *NormalImpl) SendTo(msg *chatMsg.Msg) {
	// todo
}

func (ni *NormalImpl) Pull(targetIs string) (*chatMsg.MsgPack,error) {
	// todo
	return nil,nil
}