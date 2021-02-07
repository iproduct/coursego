package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/iproduct/coursego/08-databases-lab/entities"
	"log"
)

var (
	ctx context.Context
	db  *sql.DB
)

func main() {
	db, err := sql.Open("mysql", "root:root@/golang_projects?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM projects")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var projects []entities.Project
	for rows.Next() {
		p := entities.Project{}
		if err = rows.Scan(&p.ID, &p.Name, &p.Description, &p.Budget, &p.Finished, &p.StartDate, &p.CompanyID); err != nil {
			log.Fatal(err)
		}
		projects = append(projects, p)
	}

	for i, p := range projects {
		fmt.Printf("%d -> %+v\n", i+1, p)
	}

}
