package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "user:mypassword@tcp(127.0.0.1:3306)/people?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	person := Person{
		FirstName: "Nathaniel",
		LastName:  "Payne",
	}
	// db.AutoMigrate(&Person{})
	db.Create(&person)

	fmt.Printf("%+v", db)
}

type Person struct {
	gorm.Model
	FirstName string
	LastName  string
}
