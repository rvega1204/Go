package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Connect establishes a connection to the MySQL database using environment variables.
// It loads the environment variables, constructs the Data Source Name (DSN), and attempts to connect to the database.
func Connect() (*sql.DB, error) {
	// Load environment variables from a .env file using the godotenv package.
	// This is important to load sensitive information like database credentials from a local environment.
	err := godotenv.Load()
	if err != nil {
		// If there's an error loading the .env file, return the error.
		return nil, err
	}

	// Construct the Data Source Name (DSN) string for MySQL connection.
	// The DSN format is: user:password@protocol(address)/dbname
	// Here, we get the necessary information from environment variables.
	dns := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		os.Getenv("DB_USER"),     // DB_USER is the MySQL username
		os.Getenv("DB_PASSWORD"), // DB_PASSWORD is the MySQL password
		os.Getenv("DB_HOST"),     // DB_HOST is the hostname where MySQL is running (e.g., "localhost")
		os.Getenv("DB_PORT"),     // DB_PORT is the port on which MySQL is listening (default is 3306)
		os.Getenv("DB_NAME"))     // DB_NAME is the name of the database to connect to

	// Attempt to open a connection to the MySQL database using the constructed DSN.
	// sql.Open does not establish the connection immediately, but prepares the connection for use.
	db, err := sql.Open("mysql", dns)
	if err != nil {
		// If there's an error opening the connection, return the error.
		return nil, err
	}

	// Test the connection with db.Ping().
	// db.Ping() checks if the database is reachable and that the connection works.
	if err := db.Ping(); err != nil {
		// If the database is unreachable or there is a connection issue, return the error.
		return nil, err
	}

	// Log a success message indicating the connection was successful.
	log.Println("Connection to MySQL success")

	// Return the established database connection.
	// db is an open connection to the MySQL database and will be used to execute SQL queries.
	return db, nil
}
