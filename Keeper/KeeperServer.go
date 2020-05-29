package Keeper

import (
	pb "Luka/proto"
	"Luka/util"
	"fmt"
	"golang.org/x/net/context"
	"log"
)

type Server struct {
	SubUrl string `json:"subUrl"`
	Name string `json:"name"`
}

func (s *Server) KeeperAdd(ctx context.Context, in *pb.KeeperConnectRequest) (*pb.KeeperReply, error) {
	log.Println(ctx)
	errRedis := SetKeeper(in.Name, in.KeeperUrl)
	log.Println(errRedis)
	if errRedis != nil {
		return nil,errRedis
	}
	return &pb.KeeperReply{Status: util.OK},nil
}

func (s *Server) ClientAdd(ctx context.Context, in *pb.ClientConnectRequest) (*pb.ClientReply, error) {
	log.Println(ctx)
	name, keeperUrl := GetKeeper()
	if name == "" {
		return nil ,fmt.Errorf("have no alive Keeper")
	}
	return &pb.ClientReply{
		Status:    util.OK,
		Name:	name,
		KeeperUrl: keeperUrl,
	}, nil
}