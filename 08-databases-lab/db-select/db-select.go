package main

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/iproduct/coursego/08-databases-lab/utils"
	_ "github.com/kataras/tablewriter"
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

	projects, err := utils.FindAllProjects(db)
	if err != nil {
		log.Fatal(err)
	}

	utils.PrintProjects(projects)

}
