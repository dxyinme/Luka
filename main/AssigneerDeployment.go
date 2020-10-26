package main

import (
	"flag"
	"github.com/dxyinme/LukaComm/Assigneer"
	AssigneerServer "github.com/dxyinme/LukaComm/Assigneer/Server"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	Addr = flag.String("Host", ":10197", "the assigneerServer Host")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp",*Addr)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	Assigneer.RegisterAssigneerServer(s, &AssigneerServer.Server{})
	if err = s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
