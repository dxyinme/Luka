package main

import (
	"flag"
	CynicUServer "github.com/dxyinme/LukaComm/CynicU/Server"
	"github.com/golang/glog"
)

var (
	CynicUServerAddr = flag.String("addr", ":10137", "ServerAddr")
)

func main() {
	flag.Parse()
	defer glog.Flush()
	s := &CynicUServer.Server{}
	server := s.NewCynicUServer(*CynicUServerAddr,"luka")
	if err := server.Serve(s.Lis); err != nil {
		glog.Fatalf("Server failed because of %v", err)
	}
}
