package main

import (
	"context"
	"fmt"
	"github.com/iproduct/coursego/10-modules-rest-jwtauth/rest"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Staring REST User Service ...")
	a := rest.App{}
	a.Init("root", "root", "go_rest_api")
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	a.Run(":8080")
	<- done
	log.Println("Stopping HTTP server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()
	if err := a.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %+v\n", err)
	}
	log.Print("Server exited properly")
}

