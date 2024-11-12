package handlers

import (
	"apirest/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetUsers handles the request to list all users.
// If there is an error, it sends a "Not Found" response, otherwise sends the list of users.
func GetUsers(rw http.ResponseWriter, r *http.Request) {
	// Attempt to retrieve the list of users.
	if users, err := models.ListUsers(); err != nil {
		// If an error occurs, send a "Not Found" response.
		models.SendNotFound(rw)
	} else {
		// Otherwise, send the list of users as the response.
		models.SendData(rw, users)
	}
}

// GetUser handles the request to fetch a single user by their ID.
// If the user is found, it sends the user data, otherwise sends a "Not Found" response.
func GetUser(rw http.ResponseWriter, r *http.Request) {
	// Attempt to retrieve the user based on the request's ID.
	if user, err := getUserByRequest(r); err != nil {
		// If an error occurs, send a "Not Found" response.
		models.SendNotFound(rw)
	} else {
		// Otherwise, send the user data as the response.
		models.SendData(rw, user)
	}
}

// CreateUser handles the request to create a new user.
// It decodes the user data from the request body, saves the user to the database,
// and sends the newly created user data as the response.
func CreateUser(rw http.ResponseWriter, r *http.Request) {
	// Create an empty User object to hold the incoming data.
	user := models.User{}
	// Decode the request body into the User object.
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		// If the decoding fails, send an "Unprocessable Entity" response.
		models.SendUnprocessableEntity(rw)
	} else {
		// Save the new user to the database.
		user.Save()
		// Send the newly created user data as the response.
		models.SendData(rw, user)
	}
}

// DeleteUser handles the request to delete a user by their ID.
// If the user is found, it deletes the user and sends the deleted user data as the response.
func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	// Attempt to retrieve the user based on the request's ID.
	if user, err := getUserByRequest(r); err != nil {
		// If the user is not found, send a "Not Found" response.
		models.SendNotFound(rw)
	} else {
		// Delete the user.
		user.Delete()
		// Send the deleted user data as the response.
		models.SendData(rw, user)
	}
}

// UpdateUser handles the request to update an existing user's data.
// It retrieves the user based on the ID, decodes the new user data from the request body,
// updates the user in the database, and sends the updated user data as the response.
func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	var userId int64
	// Attempt to retrieve the user based on the request's ID.
	if user, err := getUserByRequest(r); err != nil {
		// If the user is not found, send a "Not Found" response.
		models.SendNotFound(rw)
	} else {
		// Store the user's ID for later use in the update.
		userId = user.Id
	}

	// Create an empty User object to hold the updated data.
	user := models.User{}
	// Decode the new user data from the request body.
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		// If the decoding fails, send an "Unprocessable Entity" response.
		models.SendUnprocessableEntity(rw)
	} else {
		// Set the user's ID to the value retrieved earlier (preserving the original ID).
		user.Id = userId
		// Save the updated user to the database.
		user.Save()
		// Send the updated user data as the response.
		models.SendData(rw, user)
	}
}

// getUserByRequest extracts the user ID from the request and retrieves the user from the database.
// Returns the user and any error encountered during retrieval.
func getUserByRequest(r *http.Request) (models.User, error) {
	// Get the user ID from the request's URL parameters.
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])
	// Retrieve the user from the database by ID.
	if user, err := models.GetUser(userId); err != nil {
		// If an error occurs, return the error.
		return *user, err
	} else {
		// Otherwise, return the user and nil error.
		return *user, nil
	}
}
