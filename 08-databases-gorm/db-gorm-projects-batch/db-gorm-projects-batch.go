package main

import (
	"fmt"
	"github.com/iproduct/coursego/08-databases/entities"
	"github.com/iproduct/coursego/08-databases/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"time"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/gorm_projects?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Project{})
	db.AutoMigrate(&entities.Company{})

	//Get all users
	users := []entities.User{}
	result := db.Preload(clause.Associations).Find(&users) // SELECT * FROM users;
	if result.Error != nil {
		log.Fatal(result.Error) // returns error
	}
	fmt.Printf("Number of users: %d\n", result.RowsAffected) // returns found records count, equals `len(users)`
	utils.PrintUsers(users)

	//Get number of companies
	var companiesCount int64 = 0
	result = db.Model(entities.Company{}).Count(&companiesCount) // SELECT * FROM users;
	if result.Error != nil {
		log.Fatal(result.Error) // returns error
	}
	fmt.Printf("Found %d companies.\n", companiesCount)

	// Insert companies if not existing
	companies := []entities.Company{
		{Name: "Linux Foundation"},
		{Name: "Sun Microsystems"},
		{Name: "Google"},
		{Name: "Docker Inc."},
	}
	if companiesCount == 0 {
		fmt.Println("Creating sample companies:")
		db.Create(&companies)
		if result.Error != nil {
			log.Fatal(result.Error) // returns error
		}
	}

	//Get all companies
	result = db.Preload(clause.Associations).Find(&companies) // SELECT * FROM users;
	if result.Error != nil {
		log.Fatal(result.Error) // returns error
	}
	fmt.Printf("Number of companies: %d\n", result.RowsAffected) // returns found records count, equals `len(users)`
	utils.PrintCompanies(companies)

	//Get number of projects
	var projectsCount int64 = 0
	result = db.Model(entities.Project{}).Count(&projectsCount) // SELECT * FROM users;
	if result.Error != nil {
		log.Fatal(result.Error) // returns error
	}
	fmt.Printf("Found %d projects.\n", projectsCount)

	// Insert projects if not existing
	loc, _ := time.LoadLocation("Europe/Sofia")
	const shortForm = "2006-Jan-02"
	t0, _ := time.ParseInLocation(shortForm, "1991-Jan-01", loc)
	t1, _ := time.ParseInLocation(shortForm, "1996-Jan-01", loc)
	t2, _ := time.ParseInLocation(shortForm, "2009-Jan-01", loc)
	t3, _ := time.ParseInLocation(shortForm, "2013-Jan-01", loc)
	projects := []entities.Project{
		{
			Name:        "tux",
			Description: "Linux mascot project",
			Budget:      1000,
			StartDate:   t0,
			Finished:    true,
			CompanyID:   companies[0].ID,
			Users:       users,
		},
		{
			Name:        "duke",
			Description: "Java mascot project",
			Budget:      2000,
			StartDate:   t1,
			Finished:    true,
			CompanyID:   companies[1].ID,
			Users:       users,
		},
		{
			Name:        "gopher",
			Description: "Linux mascot project",
			Budget:      1000,
			StartDate:   t2,
			Finished:    true,
			CompanyID:   companies[2].ID,
			Users:       users,
		},
		{
			Name:        "moby dock",
			Description: "Docker mascot project",
			Budget:      1500,
			StartDate:   t3,
			Finished:    true,
			CompanyID:   companies[3].ID,
			Users:       users,
		},
	}
	if projectsCount == 0 {
		fmt.Println("Creating sample projects:")
		result = db.Create(&projects) // pass pointer of data to Create
		if result.Error != nil {
			log.Fatal(result.Error) // returns error
		}
		fmt.Printf("New projetcs created with IDs: ")
		for _, user := range projects {
			fmt.Printf("%d, ", user.ID)
		}
		fmt.Println()
	}

	//Get all projects
	result = db.Preload(clause.Associations).Find(&projects) // SELECT * FROM users;
	if result.Error != nil {
		log.Fatal(result.Error) // returns error
	}
	fmt.Printf("Number of projects: %d\n", result.RowsAffected) // returns found records count, equals `len(users)`
	utils.PrintProjects(projects)

	// Using association mode
	err = db.Model(&(projects[0])).Association("Users").Find(&users)
	if err != nil {
		log.Fatal(result.Error) // returns error
	}
	fmt.Printf("Users in Project '%s': %v\n", projects[0].Name, users)
}
