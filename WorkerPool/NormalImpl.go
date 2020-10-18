package WorkerPool

import (
	"context"
	"flag"
	"github.com/dxyinme/Luka/cluster/broadcast"
	"github.com/dxyinme/Luka/cluster/config"
	"github.com/dxyinme/Luka/util/syncList"
	"github.com/dxyinme/LukaComm/Assigneer"
	"github.com/dxyinme/LukaComm/chatMsg"
	"github.com/dxyinme/LukaComm/util/CoHash"
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"time"
)

const (
	// PackLimit is the max size of one pull pack
	PackLimit = 30
)

var (
	AssignHost = flag.String("assignHost", "127.0.0.1:10197", "the Assign Server Host")
)

// NormalImpl :
// **cache** the message queue of all user in this WorkerPool
// **assignToStruct** the CoHash circle of keepers
type NormalImpl struct {
	cache 	map[string]*syncList.SyncList
	assign 	CoHash.AssignToStruct
}

// Initial this WorkerPool as NormalImpl
func (ni *NormalImpl) Initial() {
	ni.cache = make(map[string]*syncList.SyncList)
	conn, err := grpc.Dial(*AssignHost, grpc.WithInsecure())
	if err != nil {
		glog.Fatal(err)
	}
	defer conn.Close()
	c := Assigneer.NewAssigneerClient(conn)
	_, err = c.AddKeeper(context.Background(), &Assigneer.AddKeeperReq{
		KeeperID: uint32(config.KeeperID),
	})
	go func() {
		for {
			select {
			case <-time.After(time.Second):
				ni.syncLocationAssignToStruct()
				break
			}
		}
	}()
}

func (ni *NormalImpl) Reduce() {

}

// SyncLocationAssignToStruct timeTask for sync the online AssignToStruct
func (ni *NormalImpl) syncLocationAssignToStruct() {
	conn, err := grpc.Dial(*AssignHost, grpc.WithInsecure())
	if err != nil {
		glog.Info(err)
	}
	defer conn.Close()
	c := Assigneer.NewAssigneerClient(conn)
	ret, err := c.SyncLocation(context.Background(), &Assigneer.SyncLocationReq{
		KeeperID: 0,
	})
	if err != nil {
		glog.Info(err)
	}
	var NewKeeperIDs []int
	for _,v := range ret.KeeperIDs {
		NewKeeperIDs = append(NewKeeperIDs, int(v))
	}
	//glog.Info(ret.KeeperIDs)
	ni.assign.SetKeeperIDs(NewKeeperIDs)
}

// SendTo in NormalImpl
// if the msg is in this keeper, send into cache,
// else redirect to the keeper it belongs to.
// user‘s position confirm by UID,
// group's position confirm by GroupID(UID in other way)
func (ni *NormalImpl) SendTo(msg *chatMsg.Msg) {
	glog.Infof("from: [%s] , target: [%s] : content: %s , Be transported.",
		msg.From, msg.Target, string(msg.Content))
	hashTarget := ni.assign.AssignTo((&CoHash.UID{Uid: msg.Target}).GetHash())
	if hashTarget == uint32(config.KeeperID) {
		if msg.MsgType == chatMsg.MsgType_Single {
			ni.sendToCacheP2P(msg)
		} else if msg.MsgType == chatMsg.MsgType_Group {
			ni.sendToCacheP2G(msg)
		} else {
			// todo
		}
	} else {
		ni.redirectMessage(msg, hashTarget)
	}
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

func (ni *NormalImpl) sendToCacheP2P(msg *chatMsg.Msg) {
	var (
		nowList *syncList.SyncList
		ok      bool
	)
	nowList, ok = ni.cache[msg.Target]
	if !ok {
		nowList = syncList.New()
		ni.cache[msg.Target] = nowList
	}
	nowList.PushBack(msg)
}

// sendToCacheP2G to which users in this group in this keeper
func (ni *NormalImpl) sendToCacheP2G(msg *chatMsg.Msg) {
	// todo
}

// redirect message to correct keeper
func (ni *NormalImpl) redirectMessage(msg *chatMsg.Msg, keeperID uint32) {
	// todo
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
