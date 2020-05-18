package Provider

import (
	pb "Luka/proto"
	"Luka/util"
	"golang.org/x/net/context"
	"log"
)

type Server struct {
	Host string
	Port string
}

// 心跳机制
func (s *Server) Confirm(ctx context.Context , in *pb.RemoteEmpty) (*pb.RemoteConfirmReply, error){
	log.Println(ctx)
	log.Println(in)
	return &pb.RemoteConfirmReply{Status: util.OK},nil
}

// 远程调用
func (s *Server) Call(ctx context.Context, in *pb.CallRequest) (*pb.CallReply, error) {
	log.Println(ctx)
	log.Println(in)
	params,err := util.TransformList(in.ParamsList,in.TypeList)
	if err != nil {
		return nil,err
	}
	resParams,resTypes := GetFuncResult(in.FuncName, params)
	return &pb.CallReply{
		Status: "OK",
		ParamsList: resParams,
		TypeList: resTypes,
	},nil
}