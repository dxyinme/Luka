package main

import (
	"context"
	"flag"
	"github.com/dxyinme/Luka/assignCli"
	"github.com/dxyinme/LukaComm/Assigneer"
	"google.golang.org/grpc"
	"log"
)

var (

	assigneerAddr = flag.String("host", "127.0.0.1:10197", "assigneer address")


	// operatorID = 1
	allKeeperInfo = flag.Bool("A", false, "get all KeeperInfo")
	// operatorID = 2
	killKeeperPro = flag.Bool("K", false, "kill such PID progress")
	keeperID = flag.String("kid", "", "keeperID of which progress you want to kill")
	// operatorID = 3
	registerNode = flag.Bool("Reg", false, "register this node as a node")
	selfIP = flag.String("ip", "", "selfIP is in cluster[in NAT]")
	selfPassword = flag.String("pwd", "", "ssh password of user[worker]")

)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*assigneerAddr, grpc.WithInsecure())
	c := Assigneer.NewAssigneerClient(conn)

	if err != nil {
		log.Println(err)
	}
	if *allKeeperInfo {
		rsp, err := c.MaintainInfo(context.Background(), &Assigneer.ClusterReq{
			OperatorID: 1,
			ReqInfo:    nil,
		})
		if err != nil {
			log.Println(err)
		} else {
			assignCli.PrintGetAllKeeperInfoRsp(rsp)
		}
	} else if *killKeeperPro {
		if *keeperID == "" {
			log.Fatal("keeperID is required")
		}
	} else if *registerNode {
		rsp, err := c.RegisterNode(context.Background(), &Assigneer.RegisterNodeReq{
			Ip: *selfIP,
			Pwd: *selfPassword,
		})
		if err != nil {
			log.Fatal(err)
		} else if rsp.RegisterInfo != "" {
			log.Fatal(rsp.RegisterInfo)
		} else {
			log.Println("register node finished")
		}

	}
}
