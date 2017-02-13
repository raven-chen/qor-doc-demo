package db

import (
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/qor/publish2"
	"github.com/qor/validations"
	"github.com/raven-chen/qor_doc_demo/config"
)

var (
	DB *gorm.DB
)

func init() {
	var err error
	var db *gorm.DB

	dbConfig := config.Config.DB
	if config.Config.DB.Adapter == "mysql" {
		db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?clientFoundRows=true&parseTime=True&loc=UTC", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name))
	} else {
		panic(errors.New("Database adapter is not supported"))
	}

	if err == nil {
		DB = db
		DB.LogMode(true)

		validations.RegisterCallbacks(DB)
		publish2.RegisterCallbacks(DB)
	} else {
		fmt.Println("can't link to db")
	}
}
