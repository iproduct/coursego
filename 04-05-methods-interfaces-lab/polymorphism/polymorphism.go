package main

import (
	"fmt"
	"strconv"
)

type Employee interface {
	GetDetails() string
}

type Manager struct {
	name        string
	age         int
	designation string
	salary      int
}

func (mgr Manager) GetDetails() string {
	return fmt.Sprintf("designation: %s, name: %s, age: %d, salary: %d",
		mgr.designation, mgr.name, mgr.age, mgr.salary)
}

type TeamLead struct {
	name     string
	age      int
	teamSize int
	salary   int
}

func (lead TeamLead) GetDetails() string {
	return fmt.Sprintf("name: %s, age: %d, salary: %d, team size: %d",
		lead.name, lead.age, lead.salary, lead.teamSize)
}

func GetAllEmployeesDetails(employes []Employee) string {
	result := ""
	for i, emp := range employes {
		result += strconv.Itoa(i+1) + ": " + emp.GetDetails() + "\n"
	}
	return result
}

func main() {
	var emp1, emp2 Employee
	emp1 = Manager{"John Smith", 48, "CEO", 6500}
	fmt.Printf("Manager 1: %s\n", emp1.GetDetails())
	emp2 = TeamLead{"Georgi Petrov", 48, 12, 4500}
	fmt.Printf("Team lead 1: %s\n", emp2.GetDetails())

	employees := []Employee{emp1, emp2}
	fmt.Printf("\nList of Employees:\n%s\n", GetAllEmployeesDetails(employees))

}
