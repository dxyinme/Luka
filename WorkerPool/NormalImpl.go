package WorkerPool

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/dxyinme/Luka/Group"
	"github.com/dxyinme/Luka/cluster/broadcast"
	"github.com/dxyinme/Luka/cluster/config"
	"github.com/dxyinme/Luka/util/ListCache"
	"github.com/dxyinme/Luka/util/syncList"
	"github.com/dxyinme/LukaComm/Assigneer"
	CynicUClient "github.com/dxyinme/LukaComm/CynicU/Client"
	"github.com/dxyinme/LukaComm/chatMsg"
	"github.com/dxyinme/LukaComm/util"
	"github.com/dxyinme/LukaComm/util/CoHash"
	"github.com/dxyinme/LukaComm/util/Const"
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

const (
	// PackLimit is the max size of one pull pack
	PackLimit = 30
)

var (
	AssignHost = flag.String("assignHost", "127.0.0.1:10197", "the Assign Server Host")

	DBServerHost = flag.String("DBServerHost", "", "DBServerHost for remove-kv")
)

// NormalImpl :
// an impl for workerPool
type NormalImpl struct {

	groupCache					map[string]*Group.Impl

	// List<*chatMsg> 	: the message queue for all users in this WorkerPool
	personCache 				*ListCache.ListCache

	// Connection during each keeper in the cluster.
	redirectClients 			map[uint32]*CynicUClient.Client

	// Lock assign and hosts.
	CoHashRWLock				sync.RWMutex

	// the CoHash circle for keepers cluster.
	assign 						CoHash.AssignToStruct

	// the hosts for keeper in cluster.
	hosts 						map[uint32]string

	// about maintain
	msgRecv						int32
	msgSend						int32
	msgNotLocal					int32

}

func (ni *NormalImpl) UseCall(in *chatMsg.UseChannel) (ret *chatMsg.UseChannel, err error) {
	ret = &chatMsg.UseChannel{
		ContentType: in.ContentType,
	}
	if in.ContentType == Const.GetGroupInfo {
		var groupNameList = make([]string ,0)
		for k,_ := range ni.groupCache {
			groupNameList = append(groupNameList, k)
		}
		ret.Content, err = util.IJson.Marshal(groupNameList)
	}
	return
}

func (ni *NormalImpl) CheckAlive(req *chatMsg.KeepAlive) (ret *chatMsg.KeepAlive) {
	ret = &chatMsg.KeepAlive{
		CheckAlive: req.CheckAlive,
	}
	ret.Errors = make([]string, 1)

	ret.MsgSend = atomic.LoadInt32(&ni.msgSend)
	atomic.AddInt32(&ni.msgSend, -ret.MsgSend)
	ret.MsgRecv = atomic.LoadInt32(&ni.msgRecv)
	atomic.AddInt32(&ni.msgRecv, -ret.MsgRecv)
	ret.MsgNotLocal = atomic.LoadInt32(&ni.msgNotLocal)
	atomic.AddInt32(&ni.msgNotLocal, -ret.MsgNotLocal)
	return
}

func (ni *NormalImpl) DeleteGroup(req *chatMsg.GroupReq) error {
	var err error
	group, ok := ni.groupCache[req.GroupName]
	if !ok || group == nil {
		err = errors.New("group has not created")
	} else {
		if group.GetMaster() != req.Uid && group.GetMaster() != "" {
			return fmt.Errorf(
				"only master can delete this group . want [%s], but [%s]",
				group.GetMaster(), req.Uid)
		}

		if !req.IsCopy {
			ni.spreadGroupOperator(Const.DeleteGroup, req)
		}
		// delete the Group
		delete(ni.groupCache, req.GroupName)
	}
	return err
}

func (ni *NormalImpl) CreateGroup(req *chatMsg.GroupReq) error {
	var err error
	group, ok := ni.groupCache[req.GroupName]
	if ok || group != nil {
		err = errors.New("group has created")
	} else {
		newGroup := Group.New(req.GroupName, req.Uid)
		if !req.IsCopy {
			err = newGroup.Join(req.Uid)
		}
		if err != nil {
			return err
		}
		ni.groupCache[req.GroupName] = newGroup
		if !req.IsCopy {
			ni.spreadGroupOperator(Const.CreateGroup, req)
		}
	}
	return err
}

func (ni *NormalImpl) LeaveGroup(req *chatMsg.GroupReq) error {
	var err error
	group, ok := ni.groupCache[req.GroupName]
	if !ok || group == nil {
		err = fmt.Errorf("no such group which's name is [%s]", req.GroupName)
	} else {
		err = group.Leave(req.Uid)
	}
	return err
}

func (ni *NormalImpl) JoinGroup(req *chatMsg.GroupReq) error {
	var err error
	group, ok := ni.groupCache[req.GroupName]
	if !ok || group == nil {
		err = fmt.Errorf("no such group which's name is [%s]", req.GroupName)
	} else {
		err = group.Join(req.Uid)
	}
	return err
}

// sync assign information
func (ni *NormalImpl) SyncLocationNotify() {
	ni.CoHashRWLock.Lock()
	defer ni.CoHashRWLock.Unlock()
	ni.SyncLocationAssignToStruct()
}

// Initial this WorkerPool as NormalImpl
func (ni *NormalImpl) Initial() {

	ni.msgNotLocal = 0
	ni.msgSend = 0
	ni.msgRecv = 0

	ni.personCache = ListCache.New()
	ni.groupCache = make(map[string]*Group.Impl)

	ni.hosts = make(map[uint32]string)
	ni.redirectClients = make(map[uint32]*CynicUClient.Client)

	conn, err := grpc.Dial(*AssignHost, grpc.WithInsecure())
	if err != nil {
		glog.Fatal(err)
	}
	defer conn.Close()
	c := Assigneer.NewAssigneerClient(conn)
	pid := strconv.Itoa(os.Getpid())
	rsp, err := c.AddKeeper(context.Background(), &Assigneer.AddKeeperReq{
		KeeperID: uint32(config.KeeperID),
		Host:     config.Host,
		Pid: 	  pid,
	})
	if err != nil {
		glog.Fatal(err)
	}
	if rsp.AckMessage != "" {
		glog.Fatal(rsp.AckMessage)
	}

	ni.SyncGroupInfo()
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
	for i := 0; i < len(ret.KeeperIDs); i++ {
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

	atomic.AddInt32(&ni.msgRecv, 1)

	if msg.MsgType == chatMsg.MsgType_Single {
		// single chat
		hashTarget := ni.assign.AssignTo((&CoHash.UID{Uid: msg.Target}).GetHash())
		if hashTarget == uint32(config.KeeperID) {
			ni.sendToCache(msg, msg.Target)
		} else {
			ni.redirectMessage(msg, hashTarget)
		}
	} else if msg.MsgType == chatMsg.MsgType_Group {
		// group chat
		if msg.Spread {
			ni.spread(msg)
		}
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
	var HostList []string
	ni.CoHashRWLock.RLock()
	for _,v := range ni.hosts {
		HostList = append(HostList, v)
	}
	ni.CoHashRWLock.RUnlock()
	err = bcForPull.Initial(&HostList)
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

// send msg[for person] to personCache
func (ni *NormalImpl) sendToCache(msg *chatMsg.Msg, target string) {
	var (
		nowList *syncList.SyncList
		ok      bool
	)
	nowList, ok = ni.personCache.Get(target)
	if !ok {
		nowList = syncList.New()
		ni.personCache.Set(msg.Target, nowList)
	}
	nowList.PushBack(msg)
}

// sendToCacheP2G to which users in this group
// ! waiting for testing
func (ni *NormalImpl) sendToCacheP2G(msg *chatMsg.Msg) {
	var group *Group.Impl
	group, ok := ni.groupCache[msg.GroupName]
	if !ok || group == nil {
		return
	}
	glog.Infof("group [%s], from [%s]", msg.GroupName, msg.From)
	group.RWmu.RLock()
	defer group.RWmu.RUnlock()
	for k,v := range group.Members {
		if v && k != msg.From {
			msgCopy := *msg
			msgCopy.Spread = false
			msgCopy.Target = k
			ni.sendToCache(&msgCopy, k)
		}
	}
}

// redirect message to keeper 'keeperID'
func (ni *NormalImpl) redirectMessage(msg *chatMsg.Msg, keeperID uint32) {
	glog.Infof("%v , keeperId : %v", msg, keeperID)

	atomic.AddInt32(&ni.msgNotLocal, 1)

	var (
		client *CynicUClient.Client
		err    error
	)
	client = ni.redirectClients[keeperID]
	if client == nil {
		client = &CynicUClient.Client{}
		err = client.Initial(ni.hosts[keeperID], time.Second*3)
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

// redirect group operators to keeper 'keeperID'
func (ni *NormalImpl) redirectGroupOperator(opNum string, req *chatMsg.GroupReq, keeperID uint32) {
	glog.Infof("%v , keeperId : %v", req, keeperID)
	var (
		client *CynicUClient.Client
		err    error
	)
	client = ni.redirectClients[keeperID]
	if client == nil {
		client = &CynicUClient.Client{}
		err = client.Initial(ni.hosts[keeperID], time.Second*3)
		if err != nil {
			glog.Error(err)
		}
		ni.redirectClients[keeperID] = client
	}
	err = client.GroupOp(opNum, req)
	if err != nil {
		glog.Error(err)
	}
}

// save msg into hard disk
func (ni *NormalImpl) saveInto(msg *chatMsg.Msg) {
	// todo
	glog.Infof("from: [%s] , target: [%s] : content : %s , Be save into hard disk.",
		msg.From, msg.Target, string(msg.Content))
}

// pullSelf in this user. strategy is FIFO
func (ni *NormalImpl) pullSelf(targetIs string) (*chatMsg.MsgPack, error) {
	var (
		nowList *syncList.SyncList
		pack    *chatMsg.MsgPack
		ok      bool
	)

	glog.Infof("Pull from : [%s]", targetIs)
	nowList, ok = ni.personCache.Get(targetIs)
	pack = &chatMsg.MsgPack{MsgList: []*chatMsg.Msg{}}
	if !ok || nowList == nil {
		return pack, nil
	}
	for len(pack.MsgList) < PackLimit {
		if nowList.Len() > 0 {
			pack.MsgList = append(pack.MsgList, nowList.Remove(nowList.Front()).(*chatMsg.Msg))
		} else {
			break
		}
	}

	atomic.AddInt32(&ni.msgSend, int32(len(pack.MsgList)))
	return pack, nil
}

// spread the group message among the cluster
func (ni *NormalImpl) spread(msg *chatMsg.Msg) {
	ni.CoHashRWLock.RLock()
	defer ni.CoHashRWLock.RUnlock()
	for k,_ := range ni.hosts {
		if k == uint32(config.KeeperID) {
			continue
		}
		msgCopy := *msg
		// set Spread false.
		msgCopy.Spread = false
		ni.redirectMessage(&msgCopy, k)
	}
}

func (ni *NormalImpl) spreadGroupOperator(opNum string, req *chatMsg.GroupReq) {
	glog.Infof("opNum [%s], group [%s] , spread", opNum, req.GroupName)
	ni.CoHashRWLock.RLock()
	defer ni.CoHashRWLock.RUnlock()
	for k,_  := range ni.hosts {
		if k == uint32(config.KeeperID) {
			continue
		}
		reqCopy := *req
		reqCopy.IsCopy = true
		ni.redirectGroupOperator(opNum, &reqCopy, k)
	}
}

// only be use in initial!
// call a remote-kv for all groupInfo
func (ni *NormalImpl) SyncGroupInfo() {

	type respType struct {
		GroupNameList []string
		UidList []string
	}
	var (
		body []byte
		respItem respType
	)

	GetUrl := "http://" + *DBServerHost + "/group/GroupInfo/GetAllGroup"
	c := http.Client{}
	resp, err := c.PostForm(GetUrl, url.Values{})
	if err != nil {
		glog.Fatal(err)
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		glog.Fatal(err)
	}
	err = util.IJson.Unmarshal(body, &respItem)
	if err != nil {
		glog.Fatal(err)
	}

	for i := 0; i < len(respItem.GroupNameList); i ++ {
		glog.Infof("groupName [%s] , Uid [%s]",
			respItem.GroupNameList[i], respItem.UidList[i])
		err = ni.CreateGroup(&chatMsg.GroupReq{
			Uid:       respItem.UidList[i],
			GroupName: respItem.GroupNameList[i],
			IsCopy:    true,
		})
		if err != nil {
			glog.Fatal(err)
		}
	}
}