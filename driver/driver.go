package driver

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {

	if DB != nil {
		return DB
	}

	var err error

	dbConfig := Config.DB
	if dbConfig.Adapter == "mssql" {
		DB, err = gorm.Open("mysql", "root:secret@tcp(localhost:3306)/db_kecamatan?parseTime=true")
		log.Println("Connected to Database Development")
	} else if dbConfig.Adapter == "postgressql" {
		DB, err = gorm.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=enable", dbConfig.UserDB, dbConfig.Password, dbConfig.Host, dbConfig.Name))
		log.Println("Connected to Database Local postgressql")
	} else if dbConfig.Adapter == "cockroachsql" {
		DB, err = gorm.Open("postgres", fmt.Sprintf("postgresql://%s@%s:%s/%s?sslmode=disable", dbConfig.UserDB, dbConfig.Host, dbConfig.Port, dbConfig.Name))
		log.Println("Connected to Database Local cockroachsql")
	}

	if err != nil {
		log.Println( "[Driver.ConnectDB] error when connect to database")
		log.Fatal(err)
	} else {
		log.Println( "SUCCES CONNECT TO DATABASE")
	}

	return DB
}
