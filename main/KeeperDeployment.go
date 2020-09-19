package main

import (
	"flag"
	"github.com/dxyinme/Luka/WorkerPool"
	ClusterConfig "github.com/dxyinme/Luka/cluster/config"
	CynicUServer "github.com/dxyinme/LukaComm/CynicU/Server"
	"github.com/golang/glog"
)

var (
	ClusterFile = flag.String("ClusterFile", "", "the file of ClusterInfo")
)

func main() {
	flag.Parse()
	defer glog.Flush()
	s := &CynicUServer.Server{}
	glog.Info(*ClusterFile)
	ClusterConfig.LoadFromFile(*ClusterFile)
	server := s.NewCynicUServer(ClusterConfig.HostAddr,"luka")
	// 先New，再bind，新的WorkerPool会被覆盖
	// bind的时候记住，务必bind初始化完成的Impl
	normalImpl := &WorkerPool.NormalImpl{}
	normalImpl.Initial()
	s.BindWorkerPool(normalImpl)
	if err := server.Serve(s.Lis); err != nil {
		glog.Fatalf("Server failed because of %v", err)
	}
}
