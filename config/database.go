package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"sever/utils"
	"strconv"
)

var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	port, err  := strconv.Atoi(utils.Env("DB_PORT"))

	if err != nil {
		fmt.Println(err)
	}

	dbConfig := DBConfig{
		Host:     utils.Env("DB_HOST"),
		Port:     port,
		User:     utils.Env("DB_USERNAME"),
		Password: utils.Env("DB_PASSWORD"),
		DBName:   utils.Env("DB_DATABASE"),
	}

	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}