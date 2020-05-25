package Keeper

import (
	pb "Luka/proto"
	"Luka/util"
	"golang.org/x/net/context"
	"log"
)

type Server struct {
	Host string `json:"host"`
	Port string `json:"port"`
	Name string `json:"name"`
}

func (s *Server) Register(ctx context.Context, in *pb.ClientConnectRequest) (*pb.RegisterReply, error) {
	log.Print(ctx)
	newKeeper := &Keeper{
		Name:     in.Name,
		IsOnline: false,
		Host:     in.Host,
		Port:     in.Port,
	}
	errRegister := SetKeeper(in.Name , newKeeper)
	reply := &pb.RegisterReply{}
	log.Println(errRegister)
	if errRegister != nil {
		reply.Status = errRegister.Error()
		return reply,errRegister
	} else {
		reply.Status = util.OK
		return reply,nil
	}

}