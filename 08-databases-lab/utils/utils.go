package utils

import (
	"database/sql"
	"github.com/iproduct/coursego/08-databases-lab/entities"
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
			return
		}
		userRows, err := db.Query("SELECT user_id FROM projects_users WHERE project_id = ?", p.ID)
		if err != nil {
			return nil, err
		}
		defer userRows.Close()
		for userRows.Next() {
			var userId uint
			if err = userRows.Scan(&userId); err != nil {
				return nil, err
			}
			p.UserID = append(p.UserID, userId)
		}
		if err = userRows.Err(); err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

func PrintProjects(projects []entities.Project) {
	printer.Print(projects)
}
