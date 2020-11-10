package WorkerPool

import (
	"context"
	"flag"
	"github.com/dxyinme/Luka/cluster/broadcast"
	"github.com/dxyinme/Luka/cluster/config"
	"github.com/dxyinme/Luka/util/syncList"
	"github.com/dxyinme/LukaComm/Assigneer"
	CynicUClient "github.com/dxyinme/LukaComm/CynicU/Client"
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
// an impl for workerPool
type NormalImpl struct {
	// List<UID>		: the UID cache for this keeper.
	groupCache		map[string]*syncList.SyncList
	// List<*chatMsg> 	: the message queue of all user in this WorkerPool
	personCache 	map[string]*syncList.SyncList

	// Connection during each keeper in the cluster.
	redirectClients map[uint32]*CynicUClient.Client

	// the CoHash circle for keepers cluster.
	assign 			CoHash.AssignToStruct

	// the hosts for keeper in cluster.
	hosts 			map[uint32]string
}

func (ni *NormalImpl) SyncLocationNotify() {
	ni.SyncLocationAssignToStruct()
}

// Initial this WorkerPool as NormalImpl
func (ni *NormalImpl) Initial() {
	ni.personCache = make(map[string]*syncList.SyncList)
	ni.groupCache = make(map[string]*syncList.SyncList)

	ni.hosts = make(map[uint32]string)
	ni.redirectClients = make(map[uint32]*CynicUClient.Client)

	conn, err := grpc.Dial(*AssignHost, grpc.WithInsecure())
	if err != nil {
		glog.Fatal(err)
	}
	defer conn.Close()
	c := Assigneer.NewAssigneerClient(conn)
	_, err = c.AddKeeper(context.Background(), &Assigneer.AddKeeperReq{
		KeeperID: 	uint32(config.KeeperID),
		Host:		config.Host,
	})
}

func (ni *NormalImpl) Reduce() {
	conn, err := grpc.Dial(*AssignHost, grpc.WithInsecure())
	if err != nil {
		glog.Fatal(err)
	}
	defer conn.Close()
	c := Assigneer.NewAssigneerClient(conn)
	_, err = c.RemoveKeeper(context.Background(), &Assigneer.RemoveKeeperReq{
		KeeperID: uint32(config.KeeperID),
	})
}

// SyncLocationAssignToStruct timeTask for sync the online AssignToStruct
func (ni *NormalImpl) SyncLocationAssignToStruct() {
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
	for i := 0 ; i < len(ret.KeeperIDs); i ++ {
		NewKeeperIDs = append(NewKeeperIDs, int(ret.KeeperIDs[i]))
		ni.hosts[ret.KeeperIDs[i]] = ret.Hosts[i]
	}
	glog.Infof("keeperIds : %v, Hosts : %v", ret.KeeperIDs, ret.Hosts)
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
	if msg.MsgType == chatMsg.MsgType_Single {
		// single chat
		hashTarget := ni.assign.AssignTo((&CoHash.UID{Uid: msg.Target}).GetHash())
		if hashTarget == uint32(config.KeeperID) {
			if msg.MsgType == chatMsg.MsgType_Single {
				ni.sendToCache(msg, msg.Target)
			}
		} else {
			ni.redirectMessage(msg, hashTarget)
		}
	} else if msg.MsgType == chatMsg.MsgType_Group {
		// group chat
		ni.sendToCacheP2G(msg)
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

func (ni *NormalImpl) sendToCache(msg *chatMsg.Msg, target string) {
	var (
		nowList *syncList.SyncList
		ok      bool
	)
	nowList, ok = ni.personCache[target]
	if !ok {
		nowList = syncList.New()
		ni.personCache[msg.Target] = nowList
	}
	nowList.PushBack(msg)
}

// sendToCacheP2G to which users in this group
// ! waiting for testing
func (ni *NormalImpl) sendToCacheP2G(msg *chatMsg.Msg) {
	nowList, ok := ni.groupCache[msg.GroupName]
	if !ok {
		return
	}

	if nowList == nil {
		return
	}
	nowList.Lock()
	for item := nowList.Front() ; item != nil ; item = item.Next() {
		UID := item.Value.(string)
		msgCopy := *msg
		msgCopy.Spread = false
		ni.sendToCache(&msgCopy, UID)
	}
	nowList.Unlock()
	if msg.Spread {
		for keeperID := range ni.hosts {
			msgCopy := *msg
			msgCopy.Spread = false
			ni.redirectMessage(&msgCopy, keeperID)
		}
	}
}

// redirect message to keeper 'keeperID'
func (ni *NormalImpl) redirectMessage(msg *chatMsg.Msg, keeperID uint32) {
	glog.Infof("%v , keeperId : %v", msg, keeperID)
	var (
		client *CynicUClient.Client
		err error
	)
	client = ni.redirectClients[keeperID]
	if client == nil {
		client = &CynicUClient.Client{}
		err = client.Initial(ni.hosts[keeperID], time.Second * 3)
		if err != nil {
			glog.Error(err)
		}
		ni.redirectClients[keeperID] = client
	}
	err = client.SendTo(msg)
	if err != nil {
		glog.Error(err)
	}
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
	nowList, ok = ni.personCache[targetIs]
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
