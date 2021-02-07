package utils

import (
	"fmt"
	"github.com/iproduct/coursego/08-databases/entities"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
)

var I2b = []bool{false, true}

// Utility functions
func PrintProjects(entities []entities.Project) {
	tableRows := []table.Row{}
	for _, p := range entities {
		row := table.Row{p.ID, p.Name, p.Description, p.Budget, p.StartDate,fmt.Sprint(p.Users)}
		tableRows = append(tableRows, row)
	}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Name", "Description", "Budget", "Start Date", "Users"})
	t.AppendRows(tableRows)
	t.Render()
}

func PrintUsers(entities []entities.User) {
	tableRows := []table.Row{}
	for _, u := range entities {
		row := table.Row{u.ID, u.FirstName,u.LastName, u.Email, u.Username, u.Password, u.Active, u.CreatedAt, u.UpdatedAt,
			fmt.Sprint(u.Projects)}
		tableRows = append(tableRows, row)
	}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "First Name", "Last Name", "Email", "Username", "Pasword", "Active", "Created", "Updated", "Projects"})
	t.AppendRows(tableRows)
	t.Render()
}

func PrintCompanies(entities []entities.Company) {
	tableRows := []table.Row{}
	for _, c := range entities {
		row := table.Row{c.ID, c.Name, fmt.Sprint(c.Projects)}
		tableRows = append(tableRows, row)
	}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Name", "Projects"})
	t.AppendRows(tableRows)
	t.Render()
}
