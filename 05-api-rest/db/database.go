package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Database connection URL
const url = "your_user:your_password@tcp(localhost:3306)/goweb_db"

// Variable to hold the database connection
var db *sql.DB

// Connect establishes a connection to the MySQL database
func Connect() {
	connection, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connection done")
	db = connection
}

// Close closes the database connection
func Close() {
	db.Close()
}

// Ping checks if the database connection is still alive
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

// CreateTable creates a new table if it doesn't already exist
func CreateTable(schema string, name string) {
	if !ExistsTable(name) {
		_, err := Exec(schema) // Executes the schema to create the table
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Table created %s", name)
	} else {
		fmt.Printf("Table %s already exists", name)
	}
}

// ExistsTable checks if a table with the given name exists in the database
func ExistsTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := Query(sql)
	if err != nil {
		fmt.Println(err)
	}

	return rows.Next()
}

// Exec is a helper function to execute SQL statements with arguments, if provided
func Exec(query string, args ...interface{}) (sql.Result, error) {
	Connect()
	result, err := db.Exec(query, args...)
	Close()
	if err != nil {
		fmt.Println(err)
	}

	return result, err
}

// Query is a helper function to execute SQL queries with arguments, if provided
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	Connect()
	rows, err := db.Query(query, args...)
	Close()
	if err != nil {
		fmt.Println(err)
	}

	return rows, err
}

// TruncateTable removes all data from the specified table
func TruncateTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE %s", tableName)
	Exec(sql)
}
