package config

import (
	"database/sql"
	"os"

	// for mysql
	_ "github.com/go-sql-driver/mysql"
)

// DB function
func DB() *sql.DB {

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	address := os.Getenv("DB_ADDRESS")
	port := os.Getenv("DB_PORT")
	_db := os.Getenv("DB")

	db, _ := sql.Open("mysql", user+":"+password+"@tcp("+address+":"+port+")/"+_db)
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
