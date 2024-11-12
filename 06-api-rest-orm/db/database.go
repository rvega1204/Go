package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DSN (Data Source Name) string to connect to the MySQL database.
// It contains the username, password, host, port, database name, and connection options.
var dsn = "user:pass@tcp(localhost:3306)/goweb_db?charset=utf8mb4&parseTime=True&loc=Local"

// Database is a function that opens a connection to the MySQL database and returns the DB instance.
// If there's an error during the connection process, it logs the error and panics.
var Database = func() (db *gorm.DB) {
	// Attempt to open the MySQL database connection using the DSN.
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		// Log the error and panic if the connection fails.
		fmt.Println("DB connection error", err)
		panic(err)
	} else {
		// Log a success message and return the DB instance if the connection is successful.
		fmt.Println("DB connection success")
		return db
	}
}()
