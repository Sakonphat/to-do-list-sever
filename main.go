package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"sever/config"
	"sever/models"
	"sever/routes"
)

func main() {

	var err error

	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))

	if err != nil {
		fmt.Println("database status: ", err)
	}

	defer config.DB.Close()


	config.DB.AutoMigrate(&models.User{})

	r := routes.SetupRouter()
	r.Run()
}