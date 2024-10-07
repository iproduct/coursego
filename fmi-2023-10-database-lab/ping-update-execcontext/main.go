package main

import (
	"context"
	"database/sql"
	"github.com/iproduct/coursego/fmi-2024-10-database-lab/entities"
	"github.com/iproduct/coursego/fmi-2024-10-database-lab/utils"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/golang_projects_2024?parseTime=true")
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

	conn, err := db.Conn(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	projects, err := GetProjects(ctx, conn)
	if err != nil {
		log.Fatal(err)
	}
	utils.PrintProjects(projects)
}

func GetProjects(ctx context.Context, conn *sql.Conn) ([]entities.Project, error) {
	rows, err := conn.QueryContext(ctx, "SELECT * FROM projects")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	projects := []entities.Project{}
	for rows.Next() {
		p := entities.Project{}
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Budget, &p.Finished, &p.StartDate, &p.CompanyID); err != nil {
			return nil, err
		}
		// TODO add User IDs to Project
		projects = append(projects, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return projects, nil
}
