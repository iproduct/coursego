package main

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/iproduct/coursego/08-databases/entities"
	"github.com/iproduct/coursego/08-databases/utils"
	_ "github.com/kataras/tablewriter"
	"log"
	"time"
)

var (
	ctx context.Context
	db  *sql.DB
)

func main() {
	db, err := sql.Open("mysql", "root:root@/golang_projects?parseTime=true")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()
	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(time.Minute * 3)

	// Ping and PingContext may be used to determine if communication with
	// the database server is still possible.
	//
	// When used in a command line application Ping may be used to establish
	// that further queries are possible; that the provided DSN is valid.
	//
	// When used in long running service Ping may be part of the health
	// checking system.
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	status := "up"
	if err := db.PingContext(ctx); err != nil {
		status = "down"
	}
	log.Println(status)

	// A *DB is a pool of connections. Call Conn to reserve a connection for exclusive use.
	conn, err := db.Conn(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close() // Return the connection to the pool.

	// Print projects before update
	projects := GetProjects(db)
	utils.PrintProjects(projects)
	//for i, proj := range projects {
	//	log.Printf("%d: %v\n", i+1, proj)
	//}

	// Update project budgets by 10% increase for project after 2020
	loc, _ := time.LoadLocation("Europe/Sofia")
	const shortForm = "2006-Jan-02"
	startDate, _ := time.ParseInLocation(shortForm, "2020-Jan-01", loc)
	result, err := conn.ExecContext(ctx, `UPDATE projects SET budget = ROUND(budget * 1.5) WHERE start_date > ?;`, startDate)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Toatal budgets updated: %d\n", rows)

	// Print projects after update
	projects = GetProjects(db)
	utils.PrintProjects(projects)
}

func GetProjects(db *sql.DB) []entities.Project {
	rows, err := db.Query("SELECT * FROM projects")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	projects := []entities.Project{}
	for rows.Next() {
		p := entities.Project{}
		if err := rows.Scan(&p.Id, &p.Name, &p.Description, &p.Budget, &p.Finished, &p.StartDate, &p.CompanyId); err != nil {
			log.Fatal(err)
		}
		userRows, err := db.Query("SELECT user_id FROM projects_users WHERE project_id = ?", p.Id)
		if err != nil {
			log.Fatal(err)
		}
		var userId int64
		for userRows.Next() {
			if err := userRows.Scan(&userId); err != nil {
				log.Fatal(err)
			}
			p.UserIds = append(p.UserIds, userId)
		}
		err = userRows.Close()
		if err != nil {
			log.Fatal(err)
		}
		if err = userRows.Err(); err != nil {
			log.Fatal(err)
		}
		projects = append(projects, p)
	}
	return projects
}
