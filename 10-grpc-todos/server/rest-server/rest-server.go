package rest_server

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/iproduct/coursego/10-grpc-todos/generated/todo_service"
	logger_grpc "github.com/iproduct/coursego/10-grpc-todos/middleware/logger-grpc"
	logger_rest "github.com/iproduct/coursego/10-grpc-todos/middleware/logger-rest"
	tracer_rest "github.com/iproduct/coursego/10-grpc-todos/middleware/tracer-rest"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// RunServer runs HTTP/REST gateway
func RunServer(ctx context.Context, grpcPort, httpPort string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := todo_service.RegisterToDoServiceHandlerFromEndpoint(ctx, mux, "localhost:"+grpcPort, opts); err != nil {
		//log.Fatalf("failed to start HTTP gateway: %v", err)
		logger_grpc.Log.Fatal("failed to start HTTP gateway", zap.String("reason", err.Error()))
	}

	srv := &http.Server{
		Addr:    ":" + httpPort,
		//Handler: mux,
		// add handler with middleware
		Handler: tracer_rest.AddRequestID(
			logger_rest.AddLogger(logger_grpc.Log, mux)),
	}

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
		}

		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()

	log.Println("starting HTTP/REST gateway...")
	return srv.ListenAndServe()
}
