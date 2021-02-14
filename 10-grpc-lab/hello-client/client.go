package main

import (
	"context"
	pb "github.com/iproduct/coursego/10-grpc-lab/helloworld"
	"google.golang.org/grpc"
	"log"
	"time"
)

const(
	address = "localhost:50051"
	defaultName = "World"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := c.SayHello(ctx, &pb.NameRequest{Name: "Trayan"})
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	log.Printf("Greeting received from hello-server: %s", resp.GetMessage())
}
