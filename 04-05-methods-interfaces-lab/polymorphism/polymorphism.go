package main

import (
	"fmt"
	"strconv"
)

func main() {
	var emp1, emp2 Employee
	emp1 = Manager{"John Smith", 48, "CEO", 6500}
	fmt.Printf("Manager 1: %s\n", emp1.GetDetails())
	emp2 = TeamLead{"Georgi Petrov", 48, 12, 4500}
	fmt.Printf("Team lead 1: %s\n", emp2.GetDetails())

	employees := []Employee{emp1, emp2}
	fmt.Printf("\nList of Employees:\n%s\n", GetAllEmployeesDetails(employees))

	dossiers := []Dossier{
		Dossier{emp1, []string{"project management", "Golang programming", "finace"}},
		Dossier{emp2,
			[]string{"Golang programming", "project management", "web developemnt", "javascript"}},
	}

	fmt.Printf("\nList of Employees:\n%s\n", GetReport(dossiers))

}
