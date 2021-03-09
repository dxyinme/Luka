package main

import (
	"flag"
	"github.com/dxyinme/Luka/WorkerPool"
	ClusterConfig "github.com/dxyinme/Luka/cluster/config"
	"github.com/dxyinme/LukaComm/CynicU/SendMsg"
	CynicUServer "github.com/dxyinme/LukaComm/CynicU/Server"
	"github.com/golang/glog"
	"os"
	"os/signal"
	"syscall"
)

var (
	// -IFC : is the config is in file
	isFileConfig = flag.Bool("IFC", false, "is file config")
	// -ICC : is commandline config
	isCmdConfig = flag.Bool("ICC", false, "is commandline config")

	//udpServerAddr = flag.String("udpAddr", ":12999", "udp sendMsg.Server listener")

	exitC = make(chan os.Signal)
	normalImpl = &WorkerPool.NormalImpl{}
)

func setUpUDPServer(w *WorkerPool.NormalImpl) {
	s := SendMsg.NewServer()

	go func() {
		if err := s.Listen(ClusterConfig.HostAddr); err != nil {
			glog.Fatal(err)
		}

	}()

	go func() {
		for {
			msgNow := s.GetMsg()
			w.SendTo(msgNow)
		}
	}()

	glog.Info("udp server set up finished")
}

func exitHandle() {
	select {
	case <- exitC:
		glog.Info("this keeper has been killed")
		normalImpl.Reduce()
		os.Exit(0)
	}
}

func main() {
	flag.Parse()
	defer glog.Flush()
	s := &CynicUServer.Server{}
	if *isFileConfig {
		glog.Info("clusterFile is : " + *ClusterConfig.ClusterFile)
		ClusterConfig.LoadFromFile()
	} else if *isCmdConfig {
		glog.Info("config is from commandline")
		ClusterConfig.LoadFromCmd()
	} else {
		glog.Fatal("no config type!")
	}
	glog.Info("listen host is " + ClusterConfig.Host)
	glog.Info("listen port is " + ClusterConfig.HostAddr)
	glog.Info("pid is ", os.Getpid())
	server := s.NewCynicUServer(ClusterConfig.HostAddr, "luka")
	// 先New，再bind，新的WorkerPool会被覆盖
	// bind的时候记住，务必bind初始化完成的Impl
	normalImpl.Initial()
	defer normalImpl.Reduce()
	s.BindWorkerPool(normalImpl)
	// set up udp server
	setUpUDPServer(normalImpl)

	signal.Notify(exitC, os.Interrupt, os.Kill, syscall.SIGTERM)
	go exitHandle()

	if err := server.Serve(s.Lis); err != nil {
		glog.Fatalf("Server failed because of %v", err)
	}
}
