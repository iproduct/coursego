package main

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/iproduct/coursego/08-databases/entities"
	"github.com/iproduct/coursego/08-databases/utils"
	"log"
	"time"
)

var (
	ctx context.Context
	db  *sql.DB
)

func main() {
	db, err := sql.Open("mysql", "root:root@/golang_projects_2021?parseTime=true")
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

	// Print projects before update
	projects := GetProjects(ctx, db)
	utils.PrintProjects(projects)

	// Update project budgets by 10% increase for project after 2020 in a single transaction
	loc, _ := time.LoadLocation("Europe/Sofia")
	const shortForm = "2006-Jan-02"
	t0, _ := time.ParseInLocation(shortForm, "1991-Jan-01", loc)
	t1, _ := time.ParseInLocation(shortForm, "1996-Jan-01", loc)
	t2, _ := time.ParseInLocation(shortForm, "2009-Jan-01", loc)
	t3, _ := time.ParseInLocation(shortForm, "2013-Jan-01", loc)
	projects = []entities.Project{
		{
			Name:        "tux2",
			Description: sql.NullString{"Linux mascot project", true},
			Budget:      1000,
			StartDate:   t0,
			Finished:    true,
			CompanyId:   2,
			UserIds:     []int64{1, 2, 3},
		},
		{
			Name:        "duke2",
			Description: sql.NullString{"Java mascot project", true},
			Budget:      2000,
			StartDate:   t1,
			Finished:    true,
			CompanyId:   2,
			UserIds:     []int64{1, 2, 3},
		},
		{
			Name:        "gopher2",
			Description: sql.NullString{"Linux mascot project", true},
			Budget:      1000,
			StartDate:   t2,
			Finished:    true,
			CompanyId:   2,
			UserIds:     []int64{1, 2, 3},
		},
		{
			Name:        "moby dock2",
			Description: sql.NullString{"Docker mascot project", true},
			Budget:      1500,
			StartDate:   t3,
			Finished:    true,
			CompanyId:   2,
			UserIds:     []int64{1, 2, 3},
		},
	}

	// BEGIN TRANSACTION
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable}) // or db.BeginTx()
	if err != nil {
		log.Println(err)
		return
	}
	// DEFER ROLLBACK
	defer tx.Rollback() // The rollback will be ignored if the tx has been committed later in the function.

	stmt, err := tx.Prepare(`INSERT INTO projects(name, description , budget, start_date, finished, company_id) VALUES( ?, ?, ?, ?, ?, ? )`)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	for i := range projects {
		projects[i].Finished = true
		result, err := stmt.Exec(projects[i].Name, projects[i].Description, projects[i].Budget, projects[i].StartDate,
			projects[i].Finished, projects[i].CompanyId)
		if err != nil {
			log.Println(err)
			return
		}
		numRows, err := result.RowsAffected()
		if err != nil || numRows != 1 {
			log.Printf("Error inserting Project: %v, %s\n", projects[i], err)
		}
		insId, err := result.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		projects[i].Id = insId
	}

	// COMMIT TRANSACTION
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	// Print projects after update
	projects = GetProjects(ctx, db)
	utils.PrintProjects(projects)
}

// Helper functions
func GetProjects(ctx context.Context, conn *sql.DB) []entities.Project {
	rows, err := conn.QueryContext(ctx, "SELECT * FROM projects")
	if err != nil {
		log.Fatal(err)
	}
	//defer rows.Close()

	var projects []entities.Project
	for rows.Next() {
		p := entities.Project{}
		//var finished []byte
		if err := rows.Scan(&p.Id, &p.Name, &p.Description, &p.Budget, &p.Finished, &p.StartDate, &p.CompanyId); err != nil {
			log.Fatal(err)
		}
		//p.Finished = utils.I2b[finished[0]]
		projects = append(projects, p)
	}

	for i := range projects {
		userRows, err := conn.QueryContext(ctx, "SELECT user_id FROM projects_users WHERE project_id = ?", projects[i].Id)
		if err != nil {
			log.Fatal(err)
		}
		var userId int64
		for userRows.Next() {
			if err := userRows.Scan(&userId); err != nil {
				log.Fatal(err)
			}
			projects[i].UserIds = append(projects[i].UserIds, userId)
		}
		err = userRows.Close()
		if err != nil {
			log.Fatal(err)
		}
		if err = userRows.Err(); err != nil {
			log.Fatal(err)
		}
	}

	return projects
}
