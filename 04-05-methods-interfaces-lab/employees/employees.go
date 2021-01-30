package employees

import (
	"fmt"
	"strconv"
)

type Employee interface {
	GetDetails() string
}

type manager struct {
	name        string
	age         int
	designation string
	salary      int
}

func NewManager(name, designation string, age, salary int) Employee {
	return manager{name, age, designation, salary}
}

func (mgr manager) GetDetails() string {
	return fmt.Sprintf("designation: %s, name: %s, age: %d, salary: %d",
		mgr.designation, mgr.name, mgr.age, mgr.salary)
}

type teamLead struct {
	name     string
	age      int
	teamSize int
	salary   int
}

func NewTeamLead(name string, teamSize, age, salary int) Employee {
	return teamLead{name, age, teamSize, salary}
}

func (lead teamLead) GetDetails() string {
	return fmt.Sprintf("name: %s, age: %d, salary: %d, team size: %d",
		lead.name, lead.age, lead.salary, lead.teamSize)
}

func GetAllEmployeesDetails(Employees []Employee) string {
	result := ""
	for i, emp := range Employees {
		result += strconv.Itoa(i+1) + ": " + emp.GetDetails() + "\n"
	}
	return result
}

type Dossier struct {
	Employee
	Qualifications []string
}

func GetReport(Dossiers []Dossier) string {
	result := ""
	for i, doss := range Dossiers {
		result += fmt.Sprintf("%d: %s -> %+v\n", i+1, doss.GetDetails(), doss.Qualifications)
	}
	return result
}
