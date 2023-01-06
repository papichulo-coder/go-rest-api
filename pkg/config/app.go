package config

import (
	"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func Connect() {
	connectionString := "postgresql://postgres:wS70rZiz2Axsf5d3jtCW@containers-us-west-131.railway.app:7335/railway"
	var err error
	db, err = gorm.Open("postgres", connectionString)

	if err != nil {
		log.Println("Could not connect to the db")
		log.Fatal(err)
	}
	log.Println("Succesfully connected to the db")
	// close connection after the call to the database connection
	defer db.Close()
}

func GetDB() *gorm.DB {
	return db
}
