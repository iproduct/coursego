package entities

import (
	"database/sql"
	"time"
)

type Project struct {
	Id          int64          `header:"ID"`
	Name        string         `header:"Name"`
	Description sql.NullString `header:"Description"`
	Budget      float64        `header:"Budget"`
	StartDate   time.Time      `header:"Start Date"`
	Finished    bool           `header:"Completed"`
	CompanyId   int64          `header:"Company ID"`
	UserIds     []int64        `header:"User IDs"`
}

type Company struct {
	Id   int64  `header:"ID"`
	Name string `header:"Name"`
}

type User struct {
	Id        int64     `header:"ID"`
	FirstName string    `header:"First Name"`
	LastName  string    `header:"Last Name"`
	Email     string    `header:"Email"`
	Username  string    `header:"Username"`
	Password  string    `header:"Password"`
	Active    bool      `header:"Active"`
	Created   time.Time `header:"Created"`
	Modified  time.Time `header:"Modified"`
}

type ProjectUser struct {
	ProjectId int64 `header:"Project ID"`
	UserId    int64 `header:"User ID"`
}
