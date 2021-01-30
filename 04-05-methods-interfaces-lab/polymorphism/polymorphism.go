package main

import (
	"fmt"
	"github.com/iproduct/coursego/04-05-methods-interfaces-lab/employees"
)

func main() {
	var emp1, emp2 employees.Employee
	emp1 = employees.NewMan{"John Smith", 48, "CEO", 6500}
	fmt.Printf("manager 1: %s\n", emp1.GetDetails())
	emp2 = employees.teamLead{"Georgi Petrov", 48, 12, 4500}
	fmt.Printf("Team lead 1: %s\n", emp2.GetDetails())

	employees := []employees.Employee{emp1, emp2}
	fmt.Printf("\nList of Employees:\n%s\n", employees.GetAllEmployeesDetails(employees))

	dossiers := []employees.Dossier{
		employees.Dossier{emp1, []string{"project management", "Golang programming", "finace"}},
		employees.Dossier{emp2,
			[]string{"Golang programming", "project management", "web developemnt", "javascript"}},
	}

	fmt.Printf("\nList of Employees:\n%s\n", employees.GetReport(dossiers))

}
