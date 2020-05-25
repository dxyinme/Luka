package client

import (
	pb "Luka/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"sync"
	"time"
)

type MyLukaClient struct {
	Name string
	Host string
	Port string
	wg   *sync.WaitGroup
}

func NewMyLukaClient(name string, host string, port string) *MyLukaClient {
	return &MyLukaClient{Name: name, Host: host, Port: port}
}

func (my *MyLukaClient) HeartBeat() {
	for {
		conn, err := grpc.Dial("127.0.0.1:6965", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		c := pb.NewRegisterClient(conn)

		ctx,cancel := context.WithTimeout(context.Background(),time.Second)

		r, errGRPC := c.Register(ctx, &pb.ClientConnectRequest{
			Name: my.Name,
			Host: my.Host,
			Port: my.Port,
		})
		if errGRPC != nil {
			log.Printf("could not register: %v", errGRPC)
			break
		}
		log.Printf("confirm: %s", r.Status)
		conn.Close()
		cancel()
		time.Sleep(10 * time.Second)
	}
}