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
	db.AutoMigrate(&entities.User{}, &entities.Project{}, &entities.Company{})

	//Get all users
	users := []entities.User{}
	result := db.Find(&users) 	// SELECT * FROM users;
	if result.Error != nil {
		log.Fatal(result.Error) // returns error
	}
	fmt.Printf("Number of users: %d\n", result.RowsAffected)// returns found records count, equals `len(users)`
	utils.PrintUsers(users)

	// Insert companies
	companies := []entities.Company {
		{Name: "Linux Foundation"},
		{Name: "Sun Microsystems"},
		{Name: "Google"},
		{Name: "Docker Inc."},
	}
	db.Create(&companies)
	//Get all companies
	result = db.Table("companies").Find(&companies) 	// SELECT * FROM users;
	if result.Error != nil {
		log.Fatal(result.Error) // returns error
	}
	fmt.Printf("Number of users: %d\n", result.RowsAffected)// returns found records count, equals `len(users)`
	utils.PrintCompanies(companies)


	// Insert projects
	loc, _ := time.LoadLocation("Europe/Sofia")
	const shortForm = "2006-Jan-02"
	t0, _ := time.ParseInLocation(shortForm,"1991-Jan-01", loc)
	t1, _ := time.ParseInLocation(shortForm,"1996-Jan-01", loc)
	t2, _ := time.ParseInLocation(shortForm,"2009-Jan-01", loc)
	t3, _ := time.ParseInLocation(shortForm,"2013-Jan-01", loc)
	projects := []entities.Project {
		{
			Name:        "tux",
			Description: "Linux mascot project",
			Budget:      1000,
			StartDate:   t0,
			Finished:    true,
			CompanyID:   companies[0].ID,
			Users:      users,
		},
		{
			Name:        "duke",
			Description: "Java mascot project",
			Budget:      2000,
			StartDate:   t1,
			Finished:    true,
			CompanyID:   companies[1].ID,
			Users:      users,
		},
		{
			Name:        "gopher",
			Description: "Linux mascot project",
			Budget:      1000,
			StartDate:   t2,
			Finished:    true,
			CompanyID:   companies[2].ID,
			Users:      users,
		},
		{
			Name:        "moby dock",
			Description: "Docker mascot project",
			Budget:      1500,
			StartDate:   t3,
			Finished:    true,
			CompanyID:   companies[3].ID,
			Users:      users,
		},
	}
	result = db.Table("projects").Create(&projects) // pass pointer of data to Create
	// batch size 100
	//db.CreateInBatches(users, 100)
	if result.Error != nil {
		log.Fatal(result.Error) // returns error
	}
	fmt.Printf("New projetcs created with IDs: ")
	for _, user := range projects {
		fmt.Printf("%d, ", user.ID)
	}
	fmt.Println()

	utils.PrintProjects(projects)

	//Get all users
	//db.Table("projects").Association("Users")
	//result = db.Joins("Users").Find(&projects)	// SELECT * FROM users;
	result = db.Preload(clause.Associations).Find(&projects)	// SELECT * FROM users with associations

	if result.Error != nil {
		log.Fatal(result.Error) // returns error
	}
	fmt.Printf("Number of all projects: %d\n", result.RowsAffected)// returns found records count, equals `len(users)`
	utils.PrintProjects(projects)

	//Print all companies again - projects should be added
	result = db.Model(&companies[0]).Table("companies").Find(&companies) 	// SELECT * FROM users;
	if result.Error != nil {
		log.Fatal(result.Error) // returns error
	}
	fmt.Printf("Number of companies: %d\n", result.RowsAffected)// returns found records count, equals `len(users)`
	utils.PrintCompanies(companies)

}