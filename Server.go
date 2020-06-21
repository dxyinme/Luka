package main

import (
	"Luka/Master"
	pb "Luka/proto"
	"Luka/util"
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"net"
)

// 用于服务器注册和请求分配的 Master-server 详情见readme
func main(){
	conf,err := util.ReadYAML("Register.yaml")
	if err != nil {
		glog.Info(err)
	}
	glog.Info(conf)
	glog.Info("hello Register!!!")
	lis, errTCP := net.Listen("tcp",conf.RegisterPort)
	if errTCP != nil {
		glog.Fatalf("failed to listen %v",conf.RegisterPort)
	}
	serverRegister := grpc.NewServer()
	pb.RegisterRegisterServer(serverRegister,&Master.Server{})
	if errGRPC := serverRegister.Serve(lis) ; errGRPC != nil {
		glog.Fatal(errGRPC)
	}
}