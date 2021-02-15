package main

import (
	"context"
	pm "github.com/iproduct/coursego/10-grpc-bidi-streaming/generated/protomath"
	"io"
	"log"
	"math/rand"

	"time"

	"google.golang.org/grpc"
)

func main() {
	rand.Seed(time.Now().Unix())

	// dail server
	conn, err := grpc.Dial(":50005", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}
	defer conn.Close()

	// create stream
	client := pm.NewMathClient(conn)
	stream, err := client.Max(context.Background())
	if err != nil {
		log.Fatalf("openn stream error %v", err)
	}

	var max int32
	done := make(chan bool)

	// first goroutine sends random increasing numbers to stream
	// and closes int after 10 iterations
	go func() {
		for i := 1; i <= 20; i++ {
			// generate random nummber and send it to stream
			rnd := int32(rand.Intn(i))
			req := pm.Request{Num: rnd}
			if err := stream.Send(&req); err != nil {
				log.Fatalf("can not send %v", err)
			}
			log.Printf("%d sent", req.Num)
			time.Sleep(time.Millisecond * 200)
		}
		if err := stream.CloseSend(); err != nil {
			log.Println(err)
		}
	}()

	// second goroutine receives data from stream
	// and saves result in max variable
	//
	// if stream is finished it closes done channel
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}
			max = resp.Result
			log.Printf("new max %d received", max)
		}
	}()

	<-done
	log.Printf("finished with max=%d", max)
}
