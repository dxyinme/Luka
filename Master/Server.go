package Master

import (
	pb "github.com/dxyinme/Luka/proto"
	"github.com/dxyinme/Luka/util"
	"github.com/golang/glog"
	"golang.org/x/net/context"
)

// todo

type Server struct {

}

// 增加Keeper
func (s *Server) KeeperAdd(ctx context.Context, in *pb.KeeperConnectRequest) (*pb.KeeperReply, error) {
	glog.Infof("new keeper added , name:[%s] url:[%s] " , in.Name, in.KeeperUrl)
	newKCh := NewKeeperChannel(in.Name, in.KeeperUrl)
	err := updateKeeper(in.Name, newKCh)
	status := util.OK
	if err != nil {
		status = util.FAIL
	}
	return &pb.KeeperReply{
		Status: status,
	},err
}

// Keeper和Server交换Client信息
func (s *Server) ClientMsgUpdate(ctx context.Context, in *pb.KeeperMsgUpdateRequest) (*pb.KeeperMsgReply, error) {
	panic("implement me")
}
