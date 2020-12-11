package assigneerServer

import (
	"context"
	"github.com/dxyinme/Luka/assigneerServer/AssignUtil"
	"github.com/dxyinme/LukaComm/Assigneer"
	CynicUClient "github.com/dxyinme/LukaComm/CynicU/Client"
	"github.com/dxyinme/LukaComm/chatMsg"
	"github.com/dxyinme/LukaComm/util/CoHash"
	"github.com/golang/glog"
	"time"
)

type Server struct {
	assignToStruct CoHash.AssignToStruct
	// KV :
	// key [ keeperID ] , value [ host ]
	hosts map[uint32]string
}

func (s *Server) Initial() {
	s.hosts = make(map[uint32]string)
}

func (s *Server) SyncLocation(context.Context, *Assigneer.SyncLocationReq) (*Assigneer.SyncLocationRsp, error) {
	var ret []uint32
	var retHost []string
	for i := 0; i < len(s.assignToStruct.KeeperIDs); i++ {
		ret = append(ret, uint32(s.assignToStruct.KeeperIDs[i]))
	}
	for i := 0; i < len(ret); i++ {
		retHost = append(retHost, s.hosts[ret[i]])
	}
	return &Assigneer.SyncLocationRsp{
		KeeperIDs: ret,
		Hosts:     retHost,
	}, nil
}

func (s *Server) RemoveKeeper(ctx context.Context, in *Assigneer.RemoveKeeperReq) (*Assigneer.AssignAck, error) {
	err := s.assignToStruct.RemoveKeeper(in.KeeperID)
	if err != nil {
		glog.Info(err)
	}
	delete(s.hosts, in.KeeperID)
	s.syncLocationNotify()
	return &Assigneer.AssignAck{
		AckMessage: "",
	}, nil
}

func (s *Server) AddKeeper(ctx context.Context, in *Assigneer.AddKeeperReq) (*Assigneer.AssignAck, error) {
	s.assignToStruct.AppendKeeper(in.KeeperID)
	s.hosts[in.KeeperID] = in.Host
	s.syncLocationNotify()
	return &Assigneer.AssignAck{
		AckMessage: "",
	}, nil
}

func (s *Server) SwitchKeeper(ctx context.Context, in *Assigneer.SwitchKeeperReq) (*Assigneer.SwitchKeeperRsp, error) {
	nowUid := &CoHash.UID{Uid: in.Uid}
	keeperID := s.assignToStruct.AssignTo(nowUid.GetHash())
	return &Assigneer.SwitchKeeperRsp{
		KeeperID: keeperID,
		Host:     s.hosts[keeperID],
	}, nil
}


// operatorID :
// 1: getAllKeeperInfo
// 2: todo
func (s *Server) MaintainInfo(ctx context.Context,in *Assigneer.ClusterReq) (ret *Assigneer.ClusterRsp,err error) {
	ret = &Assigneer.ClusterRsp{}
	if in.OperatorID == 1 {
		res := s.getAllKeeperInfo()
		ret.RspInfo = res.ToBytes()
	} else if in.OperatorID == 2 {

	}
	return
}


// get all keeper information.
func (s *Server) getAllKeeperInfo() (ret *AssignUtil.KeeperList) {
	ret = &AssignUtil.KeeperList{
		KeeperId: make([]uint32, 0),
		Lis: make([]*chatMsg.KeepAlive, 0),
	}
	for k,v := range s.hosts {
		c := CynicUClient.Client{}
		err := c.Initial(v, 3 * time.Second)
		if err != nil {
			glog.Error(err)
			continue
		}
		rsp, err := c.CheckAlive()
		if err != nil {
			glog.Error(rsp)
			continue
		}
		ret.KeeperId = append(ret.KeeperId, k)
		ret.Lis = append(ret.Lis, rsp)
	}
	return
}


func (s *Server) syncLocationNotify() {
	for _, v := range s.hosts {
		go func(host string) {
			var err error
			client := &CynicUClient.Client{}
			err = client.Initial(host, time.Second*3)
			if err != nil {
				glog.Error(err)
				return
			}
			defer client.Close()
			err = client.SyncLocationNotify()
			if err != nil {
				glog.Error(err)
			}
		}(v)
	}
}
