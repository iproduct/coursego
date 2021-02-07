package main

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/iproduct/coursego/08-databases/entities"
	_ "github.com/kataras/tablewriter"
	"github.com/lensesio/tableprinter"
	"log"
	"os"
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

	stmt, err := db.Prepare("SELECT * FROM projects")
	//"SELECT * FROM projects p JOIN projects_users pu ON p.id = pu.project_id JOIN users u on u.id = pu.user_id")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.

	rows, err := stmt.Query()
	//rows, err := db.Query(q, age)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	projects := []entities.Project{}
	for rows.Next() {
		p := entities.Project{}
		//var finished []byte
		if err := rows.Scan(&p.Id, &p.Name, &p.Description, &p.Budget, &p.Finished, &p.StartDate, &p.CompanyId); err != nil {
			log.Fatal(err)
		}
		//p.Finished = utils.I2b[finished[0]]
		log.Printf("%3d: %v\n", p.Id, p)
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

	// Print all projects formatted as table
	printer := tableprinter.New(os.Stdout)
	printer.BorderTop, printer.BorderBottom, printer.BorderLeft, printer.BorderRight = true, true, true, true
	printer.CenterSeparator = "│"
	printer.ColumnSeparator = "│"
	printer.RowSeparator = "─"
	//printer.HeaderBgColor = tablewriter.BgBlackColor
	//printer.HeaderFgColor = tablewriter.FgGreenColor

	printer.Print(projects)

	//for i, proj := range projects {
	//	log.Printf("%d: %v\n", i+1, proj)
	//}

}
