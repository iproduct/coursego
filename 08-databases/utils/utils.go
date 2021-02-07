package utils

import (
	"github.com/iproduct/coursego/08-databases/entities"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
)

var I2b = []bool{false, true}

func PrintProjects(entities []entities.Project) {
	tableRows := []table.Row{}
	for _, p := range entities {
		row := table.Row{p.Id, p.Name, p.Description, p.Budget, p.StartDate, p.UserIds}
		tableRows = append(tableRows, row)
	}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Name", "Description", "Budget", "Start Date", "User IDs"})
	t.AppendRows(tableRows)
	t.Render()
}

func PrintUsers(entities []entities.User) {
	tableRows := []table.Row{}
	for _, u := range entities {
		row := table.Row{u.Id, u.FirstName, u.LastName, u.Email, u.Username, u.Password, u.Active, u.Created, u.Modified}
		tableRows = append(tableRows, row)
	}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "First Name", "Last Name", "Email", "Username", "Pasword", "Active", "Created", "Modified"})
	t.AppendRows(tableRows)
	t.Render()
}
