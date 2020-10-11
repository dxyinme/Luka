package WorkerPool

import (
	"github.com/dxyinme/Luka/cluster/broadcast"
	"github.com/dxyinme/Luka/util/syncList"
	"github.com/dxyinme/LukaComm/chatMsg"
	"github.com/golang/glog"
)

const (
	// PackLimit is the max size of one pull pack
	PackLimit = 30
)

// NormalImpl :
// **cache** the message queue of all user in this WorkerPool
type NormalImpl struct {
	cache map[string]*syncList.SyncList
}

// Initial this WorkerPool as NormalImpl
func (ni *NormalImpl) Initial() {
	ni.cache = make(map[string]*syncList.SyncList)
}

// SendTo in NormalImpl
// if the msg is in this keeper, send into cache,
// else redirect to the keeper it belongs to.
func (ni *NormalImpl) SendTo(msg *chatMsg.Msg) {
	var (
		nowList *syncList.SyncList
		ok      bool
	)
	glog.Infof("from: [%s] , target: [%s] : content: %s , Be transported.",
		msg.From, msg.Target, string(msg.Content))
	nowList, ok = ni.cache[msg.Target]
	if !ok {
		nowList = syncList.New()
		ni.cache[msg.Target] = nowList
	}
	nowList.PushBack(msg)
	ni.saveInto(msg)
}

// Pull in NormalImpl
func (ni *NormalImpl) Pull(targetIs string) (*chatMsg.MsgPack, error) {
	return ni.pullSelf(targetIs)
}

// PullAll in NormalImpl
func (ni *NormalImpl) PullAll(targetIs string) (*chatMsg.MsgPack, error) {
	var (
		pack            = &chatMsg.MsgPack{MsgList: []*chatMsg.Msg{}}
		err       error = nil
		bcForPull *broadcast.BroadcasterForPull
	)
	pack, err = ni.pullSelf(targetIs)
	if err != nil {
		glog.Error(err)
	}
	bcForPull = &broadcast.BroadcasterForPull{}
	err = bcForPull.Initial()
	if err != nil {
		glog.Error(err)
	}
	bcForPull.SetTarget(targetIs)
	bcForPull.Do()

	for i := 0; i < bcForPull.Size(); i++ {
		respi, erri := bcForPull.GetResp(i)
		if erri == nil {
			for _, msg := range respi.MsgList {
				pack.MsgList = append(pack.MsgList, msg)
			}
		} else {
			glog.Error(erri)
		}
	}
	glog.Infof("now pull result size : %d", len(pack.MsgList))
	return pack, err
}

func (ni *NormalImpl) saveInto(msg *chatMsg.Msg) {
	glog.Infof("from: [%s] , target: [%s] : content : %s , Be save into hardware.",
		msg.From, msg.Target, string(msg.Content))
}

// pullSelf in this user. strategy is LIFO
func (ni *NormalImpl) pullSelf(targetIs string) (*chatMsg.MsgPack, error) {
	var (
		nowList *syncList.SyncList
		pack    *chatMsg.MsgPack
		ok      bool
	)
	glog.Infof("Pull from : [%s]", targetIs)
	nowList, ok = ni.cache[targetIs]
	pack = &chatMsg.MsgPack{MsgList: []*chatMsg.Msg{}}
	if !ok {
		return pack, nil
	}
	for len(pack.MsgList) < PackLimit {
		if nowList.Len() > 0 {
			pack.MsgList = append(pack.MsgList, nowList.Remove(nowList.Back()).(*chatMsg.Msg))
		} else {
			break
		}
	}
	return pack, nil
}
