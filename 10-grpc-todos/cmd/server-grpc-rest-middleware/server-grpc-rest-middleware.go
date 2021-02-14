package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"github.com/iproduct/coursego/10-grpc-todos/middleware/logger-grpc"
	"github.com/iproduct/coursego/10-grpc-todos/server/grpc-server"
	rest_server "github.com/iproduct/coursego/10-grpc-todos/server/rest-server"
	"github.com/iproduct/coursego/10-grpc-todos/service"
	"os"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// Config is configuration for Server
type Config struct {
	// gRPC grpc-server start parameters section
	// gRPC is TCP port to listen by gRPC grpc-server
	GRPCPort string

	// HTTP/REST gateway start parameters section
	// HTTPPort is TCP port to listen by HTTP/REST gateway
	HTTPPort string

	// DB Datastore parameters section
	// DatastoreDBHost is host of database
	DatastoreDBHost string
	// DatastoreDBUser is username to connect to database
	DatastoreDBUser string
	// DatastoreDBPassword password to connect to database
	DatastoreDBPassword string
	// DatastoreDBSchema is schema of database
	DatastoreDBSchema string

	// Log parameters section
	// LogLevel is global log level: Debug(-1), Info(0), Warn(1), Error(2), DPanic(3), Panic(4), Fatal(5)
	LogLevel int
	// LogTimeFormat is print time format for logger-grpc e.g. 2006-01-02T15:04:05Z07:00
	LogTimeFormat string
}

// RunServer runs gRPC grpc-server and HTTP gateway
func RunServer() error {
	ctx := context.Background()

	// get configuration
	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-server-port", "9000", "gRPC port to bind")
	flag.StringVar(&cfg.HTTPPort, "http-port", "8080", "HTTP port to bind")
	flag.StringVar(&cfg.DatastoreDBHost, "db-host", "localhost", "Database host")
	flag.StringVar(&cfg.DatastoreDBUser, "db-user", "root", "Database user")
	flag.StringVar(&cfg.DatastoreDBPassword, "db-password", "root", "Database password")
	flag.StringVar(&cfg.DatastoreDBSchema, "db-schema", "grpc-demo", "Database schema")
	flag.IntVar(&cfg.LogLevel, "log-level", -1, "Global log level")
	flag.StringVar(&cfg.LogTimeFormat, "log-time-format", "2006-01-02T15:04:05.999999999Z07:00",
		"Print time format for logger-grpc e.g. 2006-01-02T15:04:05Z07:00")
	flag.Parse()

	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC grpc-server: '%s'", cfg.GRPCPort)
	}
	if len(cfg.HTTPPort) == 0 {
		return fmt.Errorf("invalid TCP port for HTTP gateway: '%s'", cfg.HTTPPort)
	}
	// initialize logger-grpc
	if err := logger_grpc.Init(cfg.LogLevel, cfg.LogTimeFormat); err != nil {
		return fmt.Errorf("failed to initialize logger-grpc: %v", err)
	}


	// add MySQL driver specific parameter to parse date/time
	// Drop it for another database
	param := "parseTime=true"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		cfg.DatastoreDBUser,
		cfg.DatastoreDBPassword,
		cfg.DatastoreDBHost,
		cfg.DatastoreDBSchema,
		param)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}
	defer db.Close()

	API := service.NewToDoServiceServer(db)
	// run HTTP gateway
	go func() {
		_ = rest_server.RunServer(ctx, cfg.GRPCPort, cfg.HTTPPort)
	}()

	return grpc_server.RunServer(ctx, API, cfg.GRPCPort)
}

func main() {
	if err := RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}