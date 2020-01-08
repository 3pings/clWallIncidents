package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var DB *sql.DB

func init() {
	//Set Variables
	dbName := os.Getenv("dbName")
	dbPass := os.Getenv("dbPass")
	dbHost := os.Getenv("dbHost")
	dbUser := os.Getenv("dbUser")

	var err error
	//Open DB Connection
	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":3306)/" + dbName
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalln(err)
	}
}
