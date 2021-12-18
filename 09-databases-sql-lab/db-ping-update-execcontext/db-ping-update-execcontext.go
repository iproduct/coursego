package main

import (
	"09-databases-sql-lab/entities"
	"09-databases-sql-lab/utils"
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
	conn, err := db.Conn(ctx) // EXCLUSIVE USE
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close() // returns connection to connection pool
	projects, err := GetProjects(ctx, conn)
	if err != nil {
		log.Fatal(err)
	}
	utils.PrintProjects(projects)
}

func GetProjects(ctx context.Context, conn *sql.Conn) ([]entities.Project, error) {
	rows, err := conn.QueryContext(ctx, "SELECT * FROM projects")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	projects := []entities.Project{}
	for rows.Next() {
		p := entities.Project{}
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Budget,
			&p.Finished, &p.StartDate, &p.CompanyID); err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	// TODO: fillin project users
	return projects, nil
}
