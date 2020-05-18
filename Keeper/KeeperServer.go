package Keeper

import (
	pb "Luka/proto"
	"Luka/util"
	"golang.org/x/net/context"
	"log"
)

type Server struct {
	Port string
	Host string
}

// 注册
func (s *Server) Register(ctx context.Context , in *pb.RegisterRequest) (*pb.RegisterReply, error){
	log.Println(ctx)
	log.Println(in)
	err := SetService(in.Name , &Server{
		Port:          in.Port,
		Host:          in.Host,
	})
	if err != nil {
		return &pb.RegisterReply{Status: util.FAIL},nil
	}
	return &pb.RegisterReply{Status: util.OK},nil
}
// 心跳机制
func (s *Server) Confirm(ctx context.Context , in *pb.KeeperEmpty) (*pb.KeeperConfirmReply, error){
	log.Println(ctx)
	log.Println(in)
	return &pb.KeeperConfirmReply{Status: util.OK},nil
}