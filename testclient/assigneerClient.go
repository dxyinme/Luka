package main

import (
	"context"
	"github.com/dxyinme/LukaComm/Assigneer"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:10197", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := Assigneer.NewAssigneerClient(conn)
	_,err = c.AddKeeper(context.Background(), &Assigneer.AddKeeperReq{
		KeeperID: 224,
	})
	if err != nil {
		log.Fatal(err)
	}
	_,err = c.AddKeeper(context.Background(), &Assigneer.AddKeeperReq{
		KeeperID: 58,
	})
	if err != nil {
		log.Fatal(err)
	}
	ret,err := c.SyncLocation(context.Background(), &Assigneer.SyncLocationReq{
		KeeperID: 0,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(ret.KeeperIDs)
	_,err = c.RemoveKeeper(context.Background(), &Assigneer.RemoveKeeperReq{
		KeeperID: 58,
	})
	ret,err = c.SyncLocation(context.Background(), &Assigneer.SyncLocationReq{
		KeeperID: 0,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(ret.KeeperIDs)
}

