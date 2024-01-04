package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/golang_projects_2021?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(time.Minute * 3)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	status := "up"
	if err := db.PingContext(ctx); err != nil {
		status = "down"
	}

	log.Printf("Database status: %s\n", status)
}
