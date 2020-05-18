package Service

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

// 注册
func (s *Server) Register(ctx context.Context , in *pb.RegisterRequest) (*pb.RegisterReply , error){
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
func (s *Server) Confirm(ctx context.Context , in *pb.Empty) (*pb.ConfirmReply , error){
	log.Println(ctx)
	log.Println(in)
	return &pb.ConfirmReply{Status: util.OK},nil
}