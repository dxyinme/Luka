package WorkerPool

import (
	"container/list"
	"fmt"
	"github.com/dxyinme/LukaComm/chatMsg"
	"github.com/golang/glog"
)

const (
	PackLimit = 30
)
type NormalImpl struct {
	cache map[string]*list.List
}

func (ni *NormalImpl) Initial() {
	ni.cache = make(map[string]*list.List)
}

func (ni *NormalImpl) SendTo(msg *chatMsg.Msg) {
	var (
		nowList *list.List
		pack 	*chatMsg.MsgPack
		ok		bool
	)
	glog.Infof("from: %s target: %s : content: %s",msg.From, msg.Target, string(msg.Content))
	nowList,ok = ni.cache[msg.Target]
	if !ok {
		nowList = list.New()
		ni.cache[msg.Target] = nowList
	}
	if nowList.Len() == 0 {
		pack = &chatMsg.MsgPack{}
		pack.MsgList = []*chatMsg.Msg {msg}
		nowList.PushBack(pack)
	} else {
		pack = nowList.Back().Value.(*chatMsg.MsgPack)
		if len(pack.MsgList) < PackLimit {
			// 每个包的包大小上限是 ${PackLimit}
			pack.MsgList = append(pack.MsgList, msg)
		} else{
			pack = &chatMsg.MsgPack{}
			pack.MsgList = []*chatMsg.Msg {msg}
			nowList.PushBack(pack)
		}
	}

}

func (ni *NormalImpl) Pull(targetIs string) (*chatMsg.MsgPack,error) {
	var (
		nowList *list.List
		pack 	*chatMsg.MsgPack
		ok 		bool
	)
	glog.Infof("Pull from : %s",targetIs)
	nowList, ok = ni.cache[targetIs]
	if !ok {
		return nil, fmt.Errorf("NoMessage")
	}
	if nowList.Len() == 0 {
		delete(ni.cache, targetIs)
		return nil, fmt.Errorf("NoMessage")
	}
	pack = nowList.Remove(nowList.Back()).(*chatMsg.MsgPack)
	return pack,nil
}