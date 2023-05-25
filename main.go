package main

import (
	"denny/gin_gorm_practise/app/db"
	"denny/gin_gorm_practise/app/router"

	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	d := &db.Sqlite{
		DBName: "test.DB",
	}

	DB = db.InitDb(d)

	router := router.InitRouter()

	router.Run(":8081")
}
