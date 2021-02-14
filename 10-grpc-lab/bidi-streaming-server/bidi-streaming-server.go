package main

import (
	pb "github.com/iproduct/coursego/10-grpc-lab/mathmax"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedMathServer
}

func (s server) Max(srv pb.Math_MaxServer) error {
	log.Println("Staring the streaming streaming response ...")
	var max int32 = 0
	ctx := srv.Context()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// receive data from stream
		req, err := srv.Recv()
		if err == io.EOF {
			log.Println("Exit")
			return nil
		}
		if err != nil {
			log.Printf("Receive error: %v", err)
			continue
		}

		if req.GetNum() > max{
			max = req.GetNum()
			if err := srv.Send(&pb.IntResponse{Num: max}); err != nil {
				log.Printf("Send error: %v", err)
			}
			log.Printf("Sent new max: %v", max)
		}

	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMathServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve: %v", err)
	}
}
