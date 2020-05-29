package main

import (
	"Luka/Keeper"
	pb "Luka/proto"
	"Luka/util"
	"google.golang.org/grpc"
	"log"
	"net"
)

// 用于服务器注册和请求分配的 Master-server 详情见readme
func main(){
	conf,err := util.ReadYAML("Register.yaml")
	if err != nil {
		log.Println(err)
	}
	log.Println(conf)
	log.Println("hello Register!!!")
	lis, errTCP := net.Listen("tcp",conf.RegisterPort)
	if errTCP != nil {
		log.Fatalf("failed to listen %v",conf.RegisterPort)
	}
	Keeper.ResetRedis()
	//it := Keeper.SetKeeper("1", "1t")
	//if it != nil {
	//	log.Println(it)
	//}
 	//log.Println(Keeper.GetKeeper())
	serverRegister := grpc.NewServer()
	pb.RegisterRegisterServer(serverRegister,&Keeper.Server{})
	if errGRPC := serverRegister.Serve(lis) ; errGRPC != nil {
		log.Println(errGRPC)
	}
}