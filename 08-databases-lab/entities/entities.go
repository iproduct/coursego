package entities

import "time"

type Company struct {
	ID   uint
	Name string
}

type Project struct {
	ID          uint
	Name        string
	Description string
	Budget      float64
	StartDate   time.Time
	Finished    bool
	CompanyID   uint
	UserID      []uint
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
