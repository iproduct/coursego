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

	//myServer := grpc.NewServer(
	//	grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
	//		grpc_ctxtags.StreamServerInterceptor(),
	//		grpc_opentracing.StreamServerInterceptor(),
	//		grpc_prometheus.StreamServerInterceptor,
	//		grpc_zap.StreamServerInterceptor(zapLogger),
	//		grpc_auth.StreamServerInterceptor(myAuthFunction),
	//		grpc_recovery.StreamServerInterceptor(),
	//	)),
	//	grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
	//		grpc_ctxtags.UnaryServerInterceptor(),
	//		grpc_opentracing.UnaryServerInterceptor(),
	//		grpc_prometheus.UnaryServerInterceptor,
	//		grpc_zap.UnaryServerInterceptor(zapLogger),
	//		grpc_auth.UnaryServerInterceptor(myAuthFunction),
	//		grpc_recovery.UnaryServerInterceptor(),
	//	)),
	//)
	//
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
