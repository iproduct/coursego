package entities

import (
	"database/sql"
	"time"
)

type Company struct {
	ID   int64
	Name string
}

type Project struct {
	ID          int64          `header:"ID"`
	Name        string         `header:"Name"`
	Description sql.NullString `header:"Description"`
	Budget      float64        `header:"Budget"`
	StartDate   time.Time      `header:"Start Date"`
	Finished    bool           `header:"Finished"`
	CompanyID   int64          `header:"Company ID"`
	UserID      []int64        `header:"User IDs"`
}

type User struct {
	ID        int64
	FirstName string
	LastName  string
	Email     string
	Username  string
	Password  string
	Active    bool
	Created   time.Time
	Modified  time.Time
}

type ProjectUser struct {
	ProjectId int64 `header:"Project ID"`
	UserId    int64 `header:"User ID"`
}
