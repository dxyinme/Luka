package Keeper

import (
	pb "Luka/proto"
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
)

// todo 这个测试也可以删
func testAdd(position string){
	conn, errTCP := grpc.Dial(position,grpc.WithInsecure())
	if errTCP != nil {
		log.Println(errTCP)
	}
	c := pb.NewRemoteCallClient(conn)
	r, errGRPC := c.Call(context.Background(),&pb.CallRequest{
		FuncName:   "Add",
		ParamsList: []string{"1","2"},
		TypeList:   []string{"int","int"},
	})
	if errGRPC != nil {
		log.Println(errGRPC)
	}
	log.Println(r.ParamsList)
}

func CircleConfirm(){
	for i := 0 ; true ; i ++ {
		log.Println(i)
		for k,v := range FuncMap {
			conn, errTCP := grpc.Dial(v.Host + v.Port, grpc.WithInsecure())
			if errTCP != nil {
				log.Println(errTCP)
			}
			c := pb.NewRemoteCallClient(conn)
			r, errGRPC := c.Confirm(context.Background(), &pb.RemoteEmpty{})
			if errGRPC != nil {
				log.Println(errGRPC)
				continue
			}
			log.Println(k + " " + r.Status)
			go testAdd(v.Host + v.Port) //todo 删掉这个测试
			conn.Close()
		}
		time.Sleep(time.Second * 10)
	}
}