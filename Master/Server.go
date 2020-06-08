package Master

import (
	pb "Luka/proto"
	"golang.org/x/net/context"
)

// todo

type Server struct {

}


// 增加Keeper
func (s Server) KeeperAdd(ctx context.Context, in *pb.KeeperConnectRequest) (*pb.KeeperReply, error) {
	panic("implement me")
}

// Keeper和Server交换Client信息
func (s Server) ClientMsgUpdate(ctx context.Context, in *pb.KeeperMsgUpdateRequest) (*pb.KeeperMsgReply, error) {
	panic("implement me")
}