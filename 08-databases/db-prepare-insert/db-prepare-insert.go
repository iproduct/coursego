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
	db, err := sql.Open("mysql", "root:root@/golang_projects_2021?parseTime=true")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	dropDbs(db)
	createDbs(db)

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
	printer := tableprinter.New(os.Stdout)
	printer.Print(companies)

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
			Description: sql.NullString{"Linux mascot project", true},
			Budget:      1000,
			StartDate:   t0,
			Finished:    true,
			CompanyId:   companies[0].Id,
			UserIds:     []int64{1, 2, 3},
		},
		{
			Name:        "duke",
			Description: sql.NullString{"Java mascot project", true},
			Budget:      2000,
			StartDate:   t1,
			Finished:    true,
			CompanyId:   companies[1].Id,
			UserIds:     []int64{1, 2, 3},
		},
		{
			Name:        "gopher",
			Description: sql.NullString{"Linux mascot project", true},
			Budget:      1000,
			StartDate:   t2,
			Finished:    true,
			CompanyId:   companies[2].Id,
			UserIds:     []int64{1, 2, 3},
		},
		{
			Name:        "moby dock",
			Description: sql.NullString{"Docker mascot project", true},
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
	printer.Print(projects)

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

	rows, err := db.Query("SELECT * FROM projects_users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	projectsUsers := []entities.ProjectUser{}
	for rows.Next() {
		pu := entities.ProjectUser{}
		if err := rows.Scan(&pu.ProjectId, &pu.UserId); err != nil {
			log.Fatal(err)
		}
		projectsUsers = append(projectsUsers, pu)
	}
	if rows.Err() != nil {
		log.Fatal(rows.Err())
	}
	printer.Print(users)

	//	rows, err = db.Query("SELECT * FROM users")
	//	for rows.Next()
	//	PrintUsers()
}

func dropDbs(db *sql.DB) {
	res, err := db.Exec("DROP TABLE IF EXISTS `projects_users`")
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err := res.RowsAffected()
	log.Printf("'projects_users' - Rows Affected: %d %v", rowsAffected, err)

	res, err = db.Exec("DROP TABLE IF EXISTS `user_roles`")
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err = res.RowsAffected()
	log.Printf("'user_roles' - Rows Affected: %d %v", rowsAffected, err)

	res, err = db.Exec("DROP TABLE IF EXISTS `users`")
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err = res.RowsAffected()
	log.Printf("'users' - Rows Affected: %d %v", rowsAffected, err)

	res, err = db.Exec("DROP TABLE IF EXISTS `projects`")
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err = res.RowsAffected()
	log.Printf("projects - Rows Affected: %d %v", rowsAffected, err)

	res, err = db.Exec("DROP TABLE IF EXISTS `companies`")
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err = res.RowsAffected()
	log.Printf("comapnies - Rows Affected: %d %v", rowsAffected, err)

}

func createDbs(db *sql.DB) {
	res, err := db.Exec("CREATE TABLE `companies` (  `id` bigint(20) NOT NULL AUTO_INCREMENT,  `name` varchar(60) DEFAULT NULL, PRIMARY KEY (`id`)) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;")
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err := res.RowsAffected()
	log.Printf("'companies' - Rows Affected: %d %v", rowsAffected, err)

	res, err = db.Exec("CREATE TABLE `projects` ( `id` bigint(20) NOT NULL AUTO_INCREMENT, `name` varchar(60) NOT NULL, `description` varchar(1024) DEFAULT NULL, `budget` double NOT NULL, `finished` tinyint(1) NOT NULL, `start_date` date DEFAULT NULL, `company_id` bigint(20) DEFAULT NULL, PRIMARY KEY (`id`), UNIQUE KEY `name_UNIQUE` (`name`), KEY `FKrvpjk20pqyytvj6m5cutub6iq` (`company_id`), CONSTRAINT `FKrvpjk20pqyytvj6m5cutub6iq` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`)) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;")
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err = res.RowsAffected()
	log.Printf("'projects' - Rows Affected: %d %v", rowsAffected, err)

	res, err = db.Exec("CREATE TABLE `users` ( `id` bigint(20) NOT NULL AUTO_INCREMENT, `first_name` varchar(20) DEFAULT NULL, `last_name` varchar(20) DEFAULT NULL, `email` varchar(255) DEFAULT NULL, `username` varchar(30) DEFAULT NULL, `password` varchar(255) DEFAULT NULL, `active` tinyint(1) NOT NULL, `created` datetime(6) DEFAULT NULL, `modified` datetime(6) DEFAULT NULL, PRIMARY KEY (`id`), UNIQUE KEY `UK_6dotkott2kjsp8vw4d0m25fb7` (`email`), UNIQUE KEY `UK_r43af9ap4edm43mmtq01oddj6` (`username`)) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;")
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err = res.RowsAffected()
	log.Printf("'users' - Rows Affected: %d %v", rowsAffected, err)

	res, err = db.Exec("CREATE TABLE `user_roles` ( `user_id` bigint(20) NOT NULL, `role` varchar(255) DEFAULT NULL, KEY `FKhfh9dx7w3ubf1co1vdev94g3f` (`user_id`), CONSTRAINT `FKhfh9dx7w3ubf1co1vdev94g3f` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;")
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err = res.RowsAffected()
	log.Printf("user_roles - Rows Affected: %d %v", rowsAffected, err)

	res, err = db.Exec("CREATE TABLE `projects_users` ( `project_id` bigint(20) NOT NULL, `user_id` bigint(20) NOT NULL, PRIMARY KEY (`project_id`,`user_id`), KEY `FKq2sfpib7vt9mmqkmw4c9rvmca` (`user_id`), CONSTRAINT `FKq2sfpib7vt9mmqkmw4c9rvmca` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`), CONSTRAINT `FKqrnu3d0a4dnxhlpew7f7x90kh` FOREIGN KEY (`project_id`) REFERENCES `projects` (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;")
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, err = res.RowsAffected()
	log.Printf("projects_users - Rows Affected: %d %v", rowsAffected, err)

}
