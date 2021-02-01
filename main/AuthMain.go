package main

import (
	"flag"
	"github.com/dxyinme/Luka/AuthServer"
	"github.com/dxyinme/LukaComm/Auth"
	"github.com/golang/glog"
	"google.golang.org/grpc"
	"net"
)

var (
	authServerPort = flag.String("authServerPort", ":20020", "authServer listen port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", *authServerPort)
	if err != nil {
		glog.Fatal(err)
	}
	s := grpc.NewServer()
	authServer := &AuthServer.Server{}
	authServer.Initial()
	Auth.RegisterAuthServer(s, authServer)
	if err = s.Serve(lis); err != nil {
		glog.Fatal(err)
	}
}
