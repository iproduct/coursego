package entities

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type Project struct {
	gorm.Model
	Name        string    `header:"Name"`
	Description string    `header:"Description"`
	Budget      float64   `header:"Budget"`
	StartDate   time.Time `header:"Start Date"`
	Finished    bool      `header:"Completed" gorm:"'Finished' boolean"`
	CompanyID   uint      `header:"Company ID"`
	Users       []User    `gorm:"many2many:projects_users;"`
}

type Company struct {
	gorm.Model
	Name     string `header:"Name"`
	Projects []Project
}

type User struct {
	gorm.Model
	FirstName string    `header:"First Name"`
	LastName  string    `header:"Last Name"`
	Email     string    `header:"Email"`
	Username  string    `header:"Username"`
	Password  string    `header:"Password"`
	Active    bool      `gorm:"'Active' boolean"`
	Projects  []Project `gorm:"many2many:projects_users;"`
}

func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	if pw, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5); err == nil {
		tx.Statement.SetColumn("Password", "{bcrypt}"+string(pw))
	}
	return
}
