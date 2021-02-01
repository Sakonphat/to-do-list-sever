package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"sever/config"
	"sever/models"
	"sever/routes"
	"sever/services"
)

func init()  {
	services.SetupRedis()
}

func main() {

	var err error

	dbPort := config.DbURL(config.BuildDBConfig())
	fmt.Printf("DB PORT : %s \n", dbPort)

	config.DB, err = gorm.Open("mysql", dbPort)

	if err != nil {
		fmt.Println("database status: ", err)
	}

	defer config.DB.Close()


	config.DB.AutoMigrate(&models.User{})
	config.DB.AutoMigrate(&models.Task{})

	r := routes.SetupRouter()
	r.Run()
}