package main

import (
	"Luka/Provider"
	pb "Luka/proto"
	"Luka/util"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

func Add(a,b int) int {
	return a + b
}

func main(){
	conf , errYAML := util.ReadYAML("Keeper.yaml")
	if errYAML != nil {
		log.Println(errYAML)
	}
	conn, errTCP := grpc.Dial(conf.RegisterHost + conf.RegisterPort, grpc.WithInsecure())
	if errTCP != nil {
		log.Println(errTCP)
	}
	defer conn.Close()
	RegisterClient := pb.NewKeeperClient(conn)
	// Contact the server and print out its response.
	name := conf.KeeperName
	reply, errGRPC := RegisterClient.Register(context.Background(),
		&pb.RegisterRequest{
			Name:name,
			Host:conf.ServiceHost,
			Port:conf.ServicePort,
		})
	if errGRPC != nil {
		log.Println(errGRPC)
	}
	//todo 注册所有函数
	Provider.WeakUp()
	_ = Provider.AddFunc("Add", Add)
	log.Printf("Status: %s", reply.Status)
	if reply.Status == util.OK {
		sev , errTCP := net.Listen("tcp" , conf.ServicePort)
		if errTCP != nil {
			log.Println(errTCP)
		}
		s := grpc.NewServer()
		pb.RegisterRemoteCallServer(s , &Provider.Server{})
		if err := s.Serve(sev); err != nil {
			log.Println(err)
		}
	} else {
		log.Println("Register failed")
	}
}
