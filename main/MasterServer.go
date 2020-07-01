package main

import (
	"flag"
	"net"

	"github.com/dxyinme/Luka/Master"
	pb "github.com/dxyinme/Luka/proto"
	"github.com/dxyinme/Luka/util"
	"github.com/golang/glog"
	"google.golang.org/grpc"
)

func InitialMaster() {
	Master.InitialKeeperPool()
}

// 用于服务器注册和请求分配的 Master-server 详情见readme
func main() {
	InitialMaster()
	flag.Parse()
	conf, err := util.ReadYAML("Register.yaml")
	if err != nil {
		glog.Info(err)
	}
	glog.Info(conf)
	glog.Info("hello Register!!!")
	lis, errTCP := net.Listen("tcp", conf.RegisterPort)
	if errTCP != nil {
		glog.Fatalf("failed to listen %v", conf.RegisterPort)
	}
	serverRegister := grpc.NewServer()
	pb.RegisterRegisterServer(serverRegister, &Master.Server{})
	if errGRPC := serverRegister.Serve(lis); errGRPC != nil {
		glog.Fatal(errGRPC)
	}
}
