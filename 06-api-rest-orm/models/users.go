package models

import (
	"gorm/db"
)

// User struct represents a user entity in the database.
// Each User has an ID, Username, Password, and Email fields.
type User struct {
	Id       int64  `json:"id"`       // Unique identifier for the user
	Username string `json:"username"` // Username of the user
	Password string `json:"password"` // Password for the user (should be hashed in real applications)
	Email    string `json:"email"`    // Email address of the user
}

// Users type represents a collection (or list) of User entities.
type Users []User

// MigrateUser function automatically migrates the User model to the database.
// It creates or updates the 'users' table in the database based on the User struct's definition.
func MigrateUser() {
	// Call AutoMigrate to automatically create or update the User table.
	db.Database.AutoMigrate(User{})
}
