package Master

import (
	MSA "github.com/dxyinme/Luka/proto/MasterServerApi"
	"github.com/dxyinme/Luka/util"
	"github.com/golang/glog"
	"golang.org/x/net/context"
)

// todo

type Server struct {
}

func (s *Server) KeeperSync(ctx context.Context, req *MSA.KeeperSyncReq) (*MSA.KeeperSyncResp, error) {
	glog.Infof("keeperSync todo, receive %v", req.PackMsg)
	return &MSA.KeeperSyncResp{}, nil
}

// 增加Keeper
func (s *Server) KeeperAdd(ctx context.Context, req *MSA.KeeperAddReq) (*MSA.KeeperAddResp, error) {
	glog.Infof("new keeper added , name:[%s] url:[%s] ", req.Name, req.KeeperUrl)
	newKCh := NewKeeperChannel(req.Name, req.KeeperUrl)
	err := updateKeeper(req.Name, newKCh)
	status := util.OK
	if err != nil {
		status = util.FAIL
	}
	return &MSA.KeeperAddResp{
		Status: status,
	}, err
}

