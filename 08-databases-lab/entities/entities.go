package entities

import "time"

type Company struct {
	ID   uint
	Name string
}

type Project struct {
	ID          uint      `header:"ID"`
	Name        string    `header:"Name"`
	Description string    `header:"Description"`
	Budget      float64   `header:"Budget"`
	StartDate   time.Time `header:"Start Date"`
	Finished    bool      `header:"Finished"`
	CompanyID   uint      `header:"Company ID"`
	UserID      []uint    `header:"User IDs"`
}

type User struct {
	ID        uint
	FirstName string
	LastName  string
	Email     string
	Username  string
	Password  string
	Active    bool
	Created   time.Time
	Modified  time.Time
}
