package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/iproduct/coursego/08-databases/entities"
	"github.com/iproduct/coursego/08-databases/utils"
	"log"
	"time"
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

	// Update project budgets by subtracting $10000 from old projects and adding money to new projects after 01.01.2000
	startDate, _ := time.Parse("2006-Jan-01", "2000-Jan-01")
	err = UpdateProjectBudgets(ctx, db, 100, startDate)
	if err != nil {
		log.Println(err)
	}

	// Print projects after update
	projects = GetProjects(ctx, db)
	utils.PrintProjects(projects)
}

// Helper functions

// UpdateProjectBudgets updates budgets of all projects by subtracting fixed amount from each old project
// before newProjectsStart and adding it to new projects started after newProjectsStart, divided equally
func UpdateProjectBudgets(ctx context.Context, db *sql.DB, amount float64, newProjectsStart time.Time) error {

	// Get a Tx for making transaction requests.
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fail(err)
	}
	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	// Selecting all projects with start dates and budgets
	rows, err := tx.QueryContext(ctx, "SELECT id, name, start_date, budget from projects")
	if err != nil {
		return fail(fmt.Errorf("error selecting projects: %v", err))
	}
	defer rows.Close()

	type projectData struct {
		Id        int64
		Name      string
		Budget    float64
		StartDate time.Time
	}

	var results []projectData
	for rows.Next() {
		var r projectData
		err := rows.Scan(&r.Id, &r.Name, &r.StartDate, &r.Budget)
		if err != nil {
			return fail(fmt.Errorf("error scanning projects: %v", err))
		}
		results = append(results, r)
	}
	log.Printf("Project results: %v\n", results)

	sum := 0.0
	for _, pd := range results {
		if pd.StartDate.Before(newProjectsStart) {
			if pd.Budget <= amount {
				return fail(fmt.Errorf("insufficient budget=$%.2f available in project '%s'", pd.Budget, pd.Name))
			}

			// reduce old project by amount
			result, err := tx.ExecContext(ctx, `UPDATE projects SET budget = budget - ? WHERE id = ?;`, amount, pd.Id)
			if err != nil { // ROLLBACK IF ERROR
				return fail(fmt.Errorf("update failed: %v", err))
			}
			rowsAffected, err := result.RowsAffected()
			if err != nil || rowsAffected != 1 {
				return fail(fmt.Errorf("update failed: %v", err))
			}
			sum += amount
			log.Printf("Project %s budget reduced by %f.\n", pd.Name, amount)
		}
	}
	log.Printf("Total old projects budget reduction: %f\n", sum)

	// COMMIT TRANSACTION
	if err := tx.Commit(); err != nil {
		return fail(err)
	}
	return nil
}

//// Update the album inventory to remove the quantity in the order.
//_, err = tx.ExecContext(ctx, "UPDATE album SET quantity = quantity - ? WHERE id = ?",
//	quantity, albumID)
//if err != nil {
//	return fail(err)
//}
//
//// Create a new row in the album_order table.
//result, err := tx.ExecContext(ctx, "INSERT INTO album_order (album_id, cust_id, quantity, date) VALUES (?, ?, ?, ?)",
//	albumID, custID, quantity, time.Now())
//if err != nil {
//	return fail(err)
//}
//// Get the ID of the order item just created.
//orderID, err := result.LastInsertId()
//if err != nil {
//	return fail(err)
//}
//
//// Commit the transaction.
//if err = tx.Commit(); err != nil {
//	return fail(err)
//}
//
//// Return the order ID.
//return orderID, nil

// Create a helper function for preparing failure results.
func fail(err error) error {
	fmt.Errorf("UpdateProjectBudgets: %v", err)
	return err
}

func GetProjects(ctx context.Context, conn *sql.DB) []entities.Project {
	rows, err := conn.QueryContext(ctx, "SELECT * FROM projects")
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
		projects = append(projects, p)
	}
	err = rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
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
