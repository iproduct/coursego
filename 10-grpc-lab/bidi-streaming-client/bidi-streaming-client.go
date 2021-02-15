package main

import (
	"context"
	pb "github.com/iproduct/coursego/10-grpc-lab/mathmax"
	"google.golang.org/grpc"
	"io"
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	// dail server
	conn, err := grpc.Dial(":50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}
	defer conn.Close()

	// create stream
	client := pb.NewMathClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stream, err := client.Max(ctx)
	if err != nil {
		log.Fatalf("openn stream error %v", err)
	}

	var max int32

	// start receiving goroutine
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				cancel()
				return
			}
			if err != nil {
				log.Printf("can not receive: %v", err)
				return
			}
			max = resp.GetNum()
			log.Printf("new max received: %d", max)
		}
	}()

	// cancel the sending
	//go func() {
	//	<-time.After(time.Second * 2)
	//	cancel()
	//}()

	sendFor:
	for i := 1; i <= 100; i++ {
		// generate random nummber and send it to stream
		rnd := int32(rand.Intn(i))
		req := pb.IntRequest{Num: rnd}
		if err := stream.Send(&req); err != nil {
			log.Fatalf("can not send %v", err)
		}
		log.Printf("%d sent", req.Num)
		select {
		case <-time.After(time.Millisecond * 200):
		case <-stream.Context().Done():
			log.Println("sending is canceled.")
			break sendFor;
		}
	}
	if err := stream.CloseSend(); err != nil {
		log.Println(err)
	}

	log.Printf("finished with max=%d", max)
}
