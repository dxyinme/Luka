package main

import (
	"Luka/Service"
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
	sev , errTCP := net.Listen("tcp" , conf.RegisterPort)
	if errTCP != nil {
		log.Println(errTCP)
	}
	s := grpc.NewServer()
	pb.RegisterRemoteCallServer(s , &Service.Server{})
	// 心跳机制
	go Service.CircleConfirm()
	if err := s.Serve(sev); err != nil {
		log.Println(err)
	}
}