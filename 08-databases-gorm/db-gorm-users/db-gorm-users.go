package main

import (
	"fmt"
	"github.com/iproduct/coursego/08-databases/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func main() {
	//dsn := "root:root@tcp(127.0.0.1:3306)/gorm_projects?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:root@tcp(127.0.0.1:3306)/golang_projects?charset=utf8&parseTime=True&loc=Local", // data source name
		DefaultStringSize: 256, // default size for string fields
		DisableDatetimePrecision: true, // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex: true, // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn: true, // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	user := entities.User{FirstName: "Rob", LastName: "Pike2", Email: "pike2@golang.com", Username: "rob2", Password: "rob",
		Active: true, Model: gorm.Model{CreatedAt: time.Now(), UpdatedAt: time.Now()}}

	result := db.Create(&user) // pass pointer of data to Create

	if result.Error != nil {
		log.Fatal(result.Error) // returns error
	}
	fmt.Printf("New user created with ID: %d -> %+v\nRows afffected: %d\n",
		user.ID, // returns inserted data's primary key
		user,
		result.RowsAffected, // returns inserted records count
	)

}
