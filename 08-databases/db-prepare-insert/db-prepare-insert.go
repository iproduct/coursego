package main

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/iproduct/coursego/08-databases/entities"
	"github.com/lensesio/tableprinter"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
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

	// Insert companies
	companies := []entities.Company{
		{Name: "Linux Foundation"},
		{Name: "Sun Microsystems"},
		{Name: "Google"},
		{Name: "Docker Inc."},
	}
	stmt, err := db.Prepare("INSERT INTO companies(name) VALUES( ? )")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	for i, c := range companies {
		res, err := stmt.Exec(c.Name)
		if err != nil {
			log.Fatal(err)
		}
		numRows, err := res.RowsAffected()
		if err != nil || numRows != 1 {
			log.Fatal("Error inserting new Company", err)
		}
		insId, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		companies[i].Id = insId
	}

	// Insert projects
	loc, _ := time.LoadLocation("Europe/Sofia")
	const shortForm = "2006-Jan-02"
	t0, _ := time.ParseInLocation(shortForm, "1991-Jan-01", loc)
	t1, _ := time.ParseInLocation(shortForm, "1996-Jan-01", loc)
	t2, _ := time.ParseInLocation(shortForm, "2009-Jan-01", loc)
	t3, _ := time.ParseInLocation(shortForm, "2013-Jan-01", loc)
	projects := []entities.Project{
		{
			Name:        "tux",
			Description: "Linux mascot project",
			Budget:      1000,
			StartDate:   t0,
			Finished:    true,
			CompanyId:   companies[0].Id,
			UserIds:     []int64{1, 2, 3},
		},
		{
			Name:        "duke",
			Description: "Java mascot project",
			Budget:      2000,
			StartDate:   t1,
			Finished:    true,
			CompanyId:   companies[1].Id,
			UserIds:     []int64{1, 2, 3},
		},
		{
			Name:        "gopher",
			Description: "Golang mascot project",
			Budget:      1000,
			StartDate:   t2,
			Finished:    true,
			CompanyId:   companies[2].Id,
			UserIds:     []int64{1, 2, 3},
		},
		{
			Name:        "moby dock",
			Description: "Docker mascot project",
			Budget:      1500,
			StartDate:   t3,
			Finished:    true,
			CompanyId:   companies[3].Id,
			UserIds:     []int64{1, 2, 3},
		},
	}

	stmt, err = db.Prepare(`INSERT INTO projects(name, description , budget, start_date, finished, company_id) VALUES( ?, ?, ?, ?, ?, ? )`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	for i, _ := range projects {
		projects[i].Finished = true
		result, err := stmt.Exec(projects[i].Name, projects[i].Description, projects[i].Budget, projects[i].StartDate,
			projects[i].Finished, projects[i].CompanyId)
		if err != nil {
			log.Fatal(err)
		}
		numRows, err := result.RowsAffected()
		if err != nil || numRows != 1 {
			log.Fatal("Error inserting new Project", err)
		}
		insId, err := result.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		projects[i].Id = insId
	}

	// Insert users
	users := []entities.User{
		{FirstName: "Linus", LastName: "Torvalds", Email: "linus@linux.com", Username: "linus", Password: "linus"},
		{FirstName: "James", LastName: "Gosling", Email: "gosling@java.com", Username: "james", Password: "james"},
		{FirstName: "Rob", LastName: "Pike", Email: "pike@golang.com", Username: "rob", Password: "rob"},
		{FirstName: "Kamel", LastName: "Founadi", Email: "kamel@docker.com", Username: "kamel", Password: "kamel"},
	}

	stmt, err = db.Prepare(
		`INSERT INTO users(first_name, last_name, email, username, password, active, created, modified) 
		VALUES( ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	for i := range users {
		users[i].Active = true
		users[i].Created = time.Now()
		users[i].Modified = time.Now()
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(users[i].Password), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}
		users[i].Password = "{bcrypt}" + string(hashedPassword)
		result, err := stmt.Exec(users[i].FirstName, users[i].LastName, users[i].Email, users[i].Username,
			users[i].Password, users[i].Active, users[i].Created, users[i].Modified)
		if err != nil {
			log.Fatal(err)
		}
		numRows, err := result.RowsAffected()
		if err != nil || numRows != 1 {
			log.Fatal("Error inserting new User", err)
		}
		insId, err := result.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		users[i].Id = insId
	}

	printer := tableprinter.New(os.Stdout)
	printer.Print(users)

	// Connect users and projects
	stmt, err = db.Prepare(
		`INSERT INTO projects_users(project_id, user_id) VALUES( ?, ?)`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	for i := range projects {
		result, err := stmt.Exec(projects[i].Id, users[i].Id)
		if err != nil {
			log.Fatal(err)
		}
		numRows, err := result.RowsAffected()
		if err != nil || numRows != 1 {
			log.Fatal("Error inserting new relation Project_User", err)
		}
	}

	//	rows, err = db.Query("SELECT * FROM users")
	//	for rows.Next()
	//	PrintUsers()
}
