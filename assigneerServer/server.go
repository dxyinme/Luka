package assigneerServer

import (
	"context"
	"fmt"
	"github.com/dxyinme/Luka/assigneerServer/AssignUtil"
	"github.com/dxyinme/Luka/sshc"
	"github.com/dxyinme/LukaComm/Assigneer"
	CynicUClient "github.com/dxyinme/LukaComm/CynicU/Client"
	"github.com/dxyinme/LukaComm/chatMsg"
	"github.com/dxyinme/LukaComm/util/CoHash"
	"github.com/golang/glog"
	"strings"
	"time"
)

type Server struct {
	assignToStruct CoHash.AssignToStruct
	// KV :
	// key [ keeperID ] , value [ keeper information ]
	keepersInfo map[uint32]*AssignUtil.KeeperInfo
}

func (s *Server) RegisterNode(_ context.Context,in *Assigneer.RegisterNodeReq) (*Assigneer.RegisterNodeRsp, error) {
	AssignUtil.AddNode(in.Ip, in.Pwd)
	return &Assigneer.RegisterNodeRsp{
		RegisterInfo: "",
	},nil
}

func (s *Server) Initial() {
	s.keepersInfo = make(map[uint32]*AssignUtil.KeeperInfo)
}

func (s *Server) SyncLocation(context.Context, *Assigneer.SyncLocationReq) (*Assigneer.SyncLocationRsp, error) {
	var ret []uint32
	var retHost []string
	for i := 0; i < len(s.assignToStruct.KeeperIDs); i++ {
		ret = append(ret, uint32(s.assignToStruct.KeeperIDs[i]))
	}
	for i := 0; i < len(ret); i++ {
		retHost = append(retHost, s.keepersInfo[ret[i]].Host)
	}
	return &Assigneer.SyncLocationRsp{
		KeeperIDs: ret,
		Hosts:     retHost,
	}, nil
}

func (s *Server) RemoveKeeper(ctx context.Context, in *Assigneer.RemoveKeeperReq) (*Assigneer.AssignAck, error) {
	err := s.assignToStruct.RemoveKeeper(in.KeeperID)
	if err != nil {
		glog.Error(err)
	}
	err = s.sshToCloseKeeper(s.keepersInfo[in.KeeperID])
	if err != nil {
		glog.Error(err)
		return &Assigneer.AssignAck{
			AckMessage: "",
		}, err
	}
	delete(s.keepersInfo, in.KeeperID)
	s.syncLocationNotify()
	return &Assigneer.AssignAck{
		AckMessage: "",
	}, nil
}

func (s *Server) AddKeeper(ctx context.Context, in *Assigneer.AddKeeperReq) (*Assigneer.AssignAck, error) {

	ip := strings.Split(in.Host, ":")[0]
	pwd, ok := AssignUtil.Cfg.GetPassword(ip)
	glog.Infof("pwd = [%s] , host = [%v] , in.Host = [%s]", pwd, ok, in.Host)
	if !ok {
		glog.Errorf("keeperID[%d] add keeper getPassword wrong", in.KeeperID)
		return &Assigneer.AssignAck{
			AckMessage: "host has not been registered",
		}, nil
	}
	// load to assignToStruct.
	s.assignToStruct.AppendKeeper(in.KeeperID)
	s.keepersInfo[in.KeeperID] = &AssignUtil.KeeperInfo{
		Host: in.Host,
		PID: in.Pid,
		KeeperId: in.KeeperID,
	}
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
		Host:     s.keepersInfo[keeperID].Host,
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
	for k,v := range s.keepersInfo {
		c := CynicUClient.Client{}
		err := c.Initial(v.Host, 3 * time.Second)
		if err != nil {
			glog.Error(err)
			c.Close()
			continue
		}
		rsp, err := c.CheckAlive()
		if err != nil {
			glog.Error(rsp)
			c.Close()
			continue
		}
		ret.KeeperId = append(ret.KeeperId, k)
		ret.Lis = append(ret.Lis, rsp)
		c.Close()
	}
	return
}


func (s *Server) syncLocationNotify() {
	for _, v := range s.keepersInfo {
		go func(host string) {
			var err error
			client := &CynicUClient.Client{}
			defer client.Close()
			err = client.Initial(host, time.Second * 3)
			if err != nil {
				glog.Error(err)
				return
			}
			defer client.Close()
			err = client.SyncLocationNotify()
			if err != nil {
				glog.Error(err)
			}
		}(v.Host)
	}
}

func (s *Server) sshToCloseKeeper(keeper *AssignUtil.KeeperInfo) error {
	pwd,ok := AssignUtil.Cfg.GetPassword(keeper.Host)
	if !ok {
		return fmt.Errorf("no such host")
	}
	session, err := sshc.SSHConnect("worker",
		pwd, keeper.Host, 22) // ssh port.
	if err != nil {
		return err
	}
	defer session.Close()
	err = session.Run("kill " + keeper.PID)
	if err != nil {
		return err
	}
	return nil
}