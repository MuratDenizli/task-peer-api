package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB //database

func init() {

	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	// Connection stringi yaratılır
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)

	// Eğer Heroku üzerinde bir PostgreSQL'e sahipseniz, bu ayarlamaları yapmak yerine doğrudan
	// heroku tarafından verilen database url'i kullanabilirsiniz
	// dbUri := fmt.Sprintf("postgres://xxxxx@xxx.compute.amazonaws.com:5432/ddjkb1easq2mec") // Database url

	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Account{}) //Database migration
}

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
