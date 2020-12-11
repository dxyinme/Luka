package main

import (
	"context"
	"flag"
	"github.com/dxyinme/LukaComm/Assigneer"
	"google.golang.org/grpc"
	"log"
)

var (

	assigneerAddr = flag.String("host", "127.0.0.1:10197", "assigneer address")


	// operatorID = 1
	allKeeperInfo = flag.Bool("A", false, "get all KeeperInfo")

)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*assigneerAddr, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}
	if *allKeeperInfo {
		c := Assigneer.NewAssigneerClient(conn)
		rsp, err := c.MaintainInfo(context.Background(), &Assigneer.ClusterReq{
			OperatorID: 1,
			ReqInfo:    nil,
		})
		if err != nil {
			log.Println(err)
		} else {
			log.Println(rsp)
		}
	}
}
