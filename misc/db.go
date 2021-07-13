package misc

import (
	"fmt"
	"log"
	"os"

	"github.com/Waire214/newjwt/models"
	"github.com/jinzhu/gorm"
)

var conn *gorm.DB

func GetDB() *gorm.DB {
	//get database values from .env environment
	dialect := os.Getenv("DIALECT")
	host := os.Getenv("host")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")

	//database connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)
	//open connection to database
	db, err := gorm.Open(dialect, dbURI)
	//log error, if any
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}
	conn = db
	//prevent the database from closing connection
	defer db.Close()
	//make struct migration to the database - it can only be done once
	db.AutoMigrate(&models.RegistrationData{})
	return conn
}
