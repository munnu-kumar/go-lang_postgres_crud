package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Person struct {
	gorm.Model

	Name  string
	Email string `gorm:"unique_index"`
	Books []Book
}

type Book struct {
	gorm.Model

	Title      string
	Author     string
	CallNumber int `gorm:"unique_index"`
	PersonID   int
}

var (
	person = &Person{Name: "Munnu", Email: "gmail.com"}
	books  = []Book{
		{Title: "BookA", Author: "AuthorA", CallNumber: 123, PersonID: 1},
		{Title: "BookB", Author: "AuthorB", CallNumber: 1234, PersonID: 1},
	}
)

var db *gorm.DB
var err error

func main() {
	//load env properties
	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")

	// db connection string
	dbUrl := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, port)
	//open connection to db
	db, err := gorm.Open(dialect, dbUrl)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("connectted to db successfully")
	}

	//close conn when main func ends
	defer db.Close()

	//migrations to database first time activity

	db.AutoMigrate(&Person{})
	db.AutoMigrate(&Book{})
	db.Create(&person)

	for idx := range books {
		db.Create(&books[idx])
	}
}
