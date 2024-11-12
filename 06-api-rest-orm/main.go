package main

import (
	"fmt"
	"gorm/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//models.MigrateUser()

	// Initialize a new router using Gorilla Mux
	mux := mux.NewRouter()

	// Define the routes for the API and bind them to their corresponding handler functions
	// GET /api/user/ - Retrieves the list of users
	mux.HandleFunc("/api/user/", handlers.GetUsers).Methods("GET")

	// GET /api/user/{id} - Retrieves a single user by their ID
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.GetUser).Methods("GET")

	// POST /api/user/ - Creates a new user with the data in the request body
	mux.HandleFunc("/api/user/", handlers.CreateUser).Methods("POST")

	// PUT /api/user/{id} - Updates an existing user's data by their ID
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")

	// DELETE /api/user/{id} - Deletes a user by their ID
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")

	// Print a message to the console indicating the server is running
	fmt.Println("Run server: http://localhost:3000")

	// Start the server on port 3000 and log any errors that occur
	log.Fatal(http.ListenAndServe(":3000", mux))
}
