package main

import (
	"flag"
	"github.com/dxyinme/Luka/util/config"
	"net"

	"github.com/dxyinme/Luka/Master"
	MSA "github.com/dxyinme/Luka/proto/MasterServerApi"
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
	conf, err := config.ReadYAML("Register.yaml")
	if err != nil {
		glog.Info(err)
	}
	glog.Info(conf)
	glog.Info("hello Register!!!")
	lis, errTCP := net.Listen("tcp", conf.RegisterPort)
	if errTCP != nil {
		glog.Fatalf("failed to listen %v", conf.RegisterPort)
	}
	masterServer := grpc.NewServer()
	MSA.RegisterMasterServiceApiServer(masterServer, &Master.Server{})
	if errGRPC := masterServer.Serve(lis); errGRPC != nil {
		glog.Fatal(errGRPC)
	}
}
