package main

import (
	"Luka/Keeper"
	pb "Luka/proto"
	"Luka/util"
	"google.golang.org/grpc"
	"log"
	"net"
)


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
	serverRegister := grpc.NewServer()
	pb.RegisterRegisterServer(serverRegister,&Keeper.Server{})
	if errGRPC := serverRegister.Serve(lis) ; errGRPC != nil {
		log.Println(errGRPC)
	}
}