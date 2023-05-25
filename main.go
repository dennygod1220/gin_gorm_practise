package main

import (
	"denny/gin_gorm_practise/app/db"
	"denny/gin_gorm_practise/app/router"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	var d db.IDb

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if os.Getenv("APP_ENV") == "testing" {
		d = &db.Sqlite{}
	} else {
		d = &db.Mysql{}
	}

	DB = db.InitDb(d.NewConnInfo())

	router := router.InitRouter()

	router.Run(":8081")
}
