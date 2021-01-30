package main

import (
	"fmt"
	"strconv"
)

type Employee interface {
	GetDetails() string
}

type Manager struct {
	Name        string
	Age         int
	Designation string
	Salary      int
}

func (mgr Manager) GetDetails() string {
	return mgr.Name + " " + strconv.Itoa(mgr.Age)
}

type TeamLead struct {
	Name     string
	Age      int
	TeamSize string
	Salary   int
}

func (ld TeamLead) GetDetails() string {
	return ld.Name + " " + strconv.Itoa(ld.Age)
}

// Functions taking interface type parameter -> runtime polymorphism
func GetUserDetails(emp Employee) string {
	return emp.GetDetails()
}

func GetDetailsForAllEmployes(employees []Employee) string {
	result := ""
	for i, emp := range employees {
		result += strconv.Itoa(i+1) + ": " + emp.GetDetails() + "\n"
	}
	return result
}

type Dossier struct {
	employee       Employee
	qualifications []string
}

func main() {
	// create new Manager
	manager := Manager{Name: "Hristo Atanasov", Age: 45, Designation: "CEO", Salary: 6500}

	// creating new TeamLead
	teamLead := TeamLead{Name: "George Petrov", Age: 35, TeamSize: "12", Salary: 4500}

	// interface typed valiable -> different dynamic types can be assigned
	var empInterface Employee

	//Manager Object assigned to Interface type since Interface Contract is satisfied
	empInterface = manager
	fmt.Printf("%T: %v\n", empInterface, GetUserDetails(empInterface))

	// Interface can be used to invoke function of either Lead or Manager...
	empInterface = teamLead
	fmt.Printf("%T: %v\n", empInterface, GetUserDetails(empInterface))

	allEmployees := []Employee{manager, teamLead}

	fmt.Printf("\nAll Employees:\n%s", GetDetailsForAllEmployes(allEmployees))

	dossier1 := Dossier{
		Manager{Name: "John Smith", Age: 42, Designation: "Developer", Salary: 10},
		[]string{"Golang", "JavaScript"},
	}
	fmt.Printf("\nDossier: %+v\n", dossier1)
}
