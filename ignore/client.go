// +build ignore

package main

import (
	pb "Luka/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

const (
	host = "127.0.0.1:6965"
)
func main(){
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewRegisterClient(conn)
	// Contact the server and print out its response.
	name := "test"
	r, errGRPC := c.Register(context.Background(), &pb.RegisterRequest{
		Name: name,
		Host: "127.0.0.1",
		Port: ":9966",
	})
	if errGRPC != nil {
		log.Fatalf("could not register: %v", errGRPC)
	}
	log.Printf("register message : %s", r.Status)
}

