package handlers

import (
	"encoding/json"
	"gorm/db"
	"gorm/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// GetUsers handles the request to list all users from the database.
// It fetches the list of users, and if successful, sends the list of users in the response with a 200 OK status.
// If an error occurs, it sends a "Not Found" response.
func GetUsers(rw http.ResponseWriter, r *http.Request) {
	users := models.Users{}
	// Fetch all users from the database.
	db.Database.Find(&users)
	// Send the list of users in the response.
	sendData(rw, users, http.StatusOK)
}

// GetUser handles the request to fetch a single user by their ID.
// It extracts the user ID from the request, fetches the user, and sends the user data in the response.
// If the user is not found, it sends a "Not Found" response.
func GetUser(rw http.ResponseWriter, r *http.Request) {
	// Try to retrieve the user by their ID from the request.
	if user, err := getUserByID(r); err != nil {
		// If user not found, send an error response.
		sendError(rw, http.StatusNotFound)
	} else {
		// Send the found user data in the response.
		sendData(rw, user, http.StatusOK)
	}
}

// getUserByID extracts the user ID from the request and retrieves the corresponding user from the database.
// It returns the user and any error encountered during the retrieval process.
func getUserByID(r *http.Request) (models.User, *gorm.DB) {
	vars := mux.Vars(r)
	// Extract user ID from the URL path parameter.
	userId, _ := strconv.Atoi(vars["id"])

	user := models.User{}
	// Attempt to fetch the user by the ID.
	if err := db.Database.First(&user, userId); err.Error != nil {
		// Return the error if the user is not found.
		return user, err
	} else {
		// Return the user and nil error if the user is found.
		return user, nil
	}
}

// CreateUser handles the request to create a new user.
// It decodes the user data from the request body, saves the user to the database,
// and sends the newly created user data in the response with a 201 Created status.
func CreateUser(rw http.ResponseWriter, r *http.Request) {
	// Create an empty User object to hold the incoming data.
	user := models.User{}
	// Decode the incoming JSON request body into the User object.
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		// If decoding fails, send an error response with status 422 Unprocessable Entity.
		sendError(rw, http.StatusUnprocessableEntity)
	} else {
		// Save the new user to the database.
		db.Database.Save(&user)
		// Send the newly created user in the response with a 201 Created status.
		sendData(rw, user, http.StatusCreated)
	}
}

// DeleteUser handles the request to delete a user by their ID.
// It attempts to fetch the user from the database, deletes it, and sends the deleted user's data in the response.
// If the user is not found, it sends a "Not Found" response.
func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	// Attempt to retrieve the user by ID from the request.
	if user, err := getUserByID(r); err != nil {
		// If user not found, send an error response.
		sendError(rw, http.StatusNotFound)
	} else {
		// Delete the user from the database.
		db.Database.Delete(&user)
		// Send the deleted user data in the response.
		sendData(rw, user, http.StatusOK)
	}
}

// UpdateUser handles the request to update an existing user's data.
// It retrieves the user by their ID, decodes the new user data from the request body,
// updates the user in the database, and sends the updated user data in the response.
func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	var userId int64
	// Try to retrieve the current user by ID from the request.
	if user_ant, err := getUserByID(r); err != nil {
		// If user not found, send an error response.
		sendError(rw, http.StatusNotFound)
	} else {
		// Save the current user's ID to update the correct user.
		userId = user_ant.Id
		user := models.User{}
		// Decode the new user data from the request body.
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&user); err != nil {
			// If decoding fails, send an error response with status 422 Unprocessable Entity.
			sendError(rw, http.StatusUnprocessableEntity)
		} else {
			// Assign the original user ID to the updated user to avoid overwriting it.
			user.Id = userId
			// Save the updated user to the database.
			db.Database.Save(&user)
			// Send the updated user data in the response.
			sendData(rw, user, http.StatusOK)
		}
	}
}
