package main

import (
	"flag"
	"github.com/dxyinme/Luka/assigneerServer"
	"github.com/dxyinme/LukaComm/Assigneer"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	Addr = flag.String("Host", ":10197", "the assigneerServer Host")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", *Addr)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	assServer := &assigneerServer.Server{}
	assServer.Initial()
	Assigneer.RegisterAssigneerServer(s, assServer)
	if err = s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
