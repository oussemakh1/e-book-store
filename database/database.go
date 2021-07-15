package database

import (
	"fmt"
	"log"
	"os"

	// imported for unknown drive issue
	"rest-api/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func GetDB(service string) (DB *gorm.DB, err error) {

	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	fmt.Println("mysqlUser: ", mysqlUser)
	fmt.Println("mysqlPassword: ", mysqlPassword)
	fmt.Println("mysqlDatabase: ", mysqlDatabase)

	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(mysql_store:3306)/%s?charset=utf8&"+
		"parseTime=True&loc=Local", mysqlUser, mysqlPassword, mysqlDatabase))
	if err != nil {
		log.Println("!!! ERROR IN DB : ", err)
		return DB, err
	}

	db.DB().SetMaxIdleConns(10)

	if service == "DEV" {
		db.LogMode(true)
	}

	DB = db

	return DB, nil
}

func Migrate() error {
	//Connect to database
	db, err := GetDB("API")
	defer db.Close()
	if err != nil {
		return err
	}
	//Migrate tables
	db.AutoMigrate(&models.Book{})
	return nil
}
