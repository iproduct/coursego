package employees

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

type Dossier struct {
	Employee
	qualifications []string
}

func GetReport(dossiers []Dossier) string {
	result := ""
	for i, doss := range dossiers {
		result += fmt.Sprintf("%d: %s -> %+v\n", i+1, doss.GetDetails(), doss.qualifications)
	}
	return result
}
