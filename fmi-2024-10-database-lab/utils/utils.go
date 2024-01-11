package utils

import (
	"database/sql"
	"github.com/iproduct/coursego/fmi-2024-10-database-lab/entities"
	"github.com/lensesio/tableprinter"
	"os"
)

var printer = tableprinter.New(os.Stdout)

func init() {
	printer.BorderLeft, printer.BorderRight, printer.BorderTop, printer.BorderBottom = true, true, true, true
	printer.CenterSeparator = "+"
	printer.ColumnSeparator = "|"
}

func FindAllProjects(db *sql.DB) (projects []entities.Project, err error) {
	stmt, err := db.Prepare("SELECT * FROM projects")
	if err != nil {
		return
	}
	defer stmt.Close()
	//rows, err := db.Query("SELECT * FROM projects")

	rows, err := stmt.Query()
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		p := entities.Project{}
		if err = rows.Scan(&p.ID, &p.Name, &p.Description, &p.Budget, &p.Finished, &p.StartDate, &p.CompanyID); err != nil {
			return nil, err
		}
		userRows, err2 := db.Query("SELECT user_id FROM projects_users WHERE project_id = ?", p.ID)
		if err2 != nil {
			return nil, err2
		}
		// defer userRows.Close() - defer in a loop is anti-pattern
		for userRows.Next() {
			var userId int64
			if err = userRows.Scan(&userId); err != nil {
				return
			}
			p.UserID = append(p.UserID, userId)
		}
		// If the database is being written to ensure to check for Close
		// errors that may be returned from the driver. The query may
		// encounter an auto-commit error and be forced to rollback changes.
		err = userRows.Close()
		if err != nil {
			return
		}

		// Rows.Err will report the last error encountered by Rows.Scan.
		if err = userRows.Err(); err != nil {
			return
		}
		projects = append(projects, p)
	}
	// If the database is being written to ensure to check for Close
	// errors that may be returned from the driver. The query may
	// encounter an auto-commit error and be forced to rollback changes.
	err = rows.Close()
	if err != nil {
		return
	}

	// Rows.Err will report the last error encountered by Rows.Scan.
	if err = rows.Err(); err != nil {
		return
	}
	return
}

func PrintProjects(projects []entities.Project) {
	printer.Print(projects)
}
