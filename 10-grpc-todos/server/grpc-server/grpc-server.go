package grpc_server

import (
	"context"
	"github.com/iproduct/coursego/10-grpc-todos/generated/todo_service"
	"github.com/iproduct/coursego/10-grpc-todos/middleware"
	"github.com/iproduct/coursego/10-grpc-todos/middleware/logger-grpc"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

// RunServer runs gRPC service to publish ToDo service
func RunServer(ctx context.Context, API todo_service.ToDoServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// gRPC server statup options
	opts := []grpc.ServerOption{}

	// add middleware
	opts = middleware.AddLogging(logger_grpc.Log, opts)

	// register service
	server := grpc.NewServer(opts...)
	todo_service.RegisterToDoServiceServer(server, API)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC grpc-server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC grpc-server
	log.Println("starting gRPC grpc-server...")
	return server.Serve(listen)
}