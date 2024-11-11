package models

import (
	"fmt"
	"gomysql/db"
)

// User struct represents a user in the database
type User struct {
	Id       int64
	Username string
	Password string
	Email    string
}

// Users type represents a list of User
type Users []User

// UserSchema defines the SQL statement to create the "users" table
const UserSchema = `CREATE TABLE users (
	id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(30) NOT NULL,
	password VARCHAR(100) NOT NULL,
	email VARCHAR(50),
	create_data TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`

// NewUser creates and returns a new User instance with the provided details
func NewUser(username, password, email string) *User {
	user := &User{Username: username, Password: password, Email: email}
	return user
}

// Private method to insert a new user into the database
func (user *User) insert() {
	sql := "INSERT users SET username=?, password=?, email=?"
	result, _ := db.Exec(sql, user.Username, user.Password, user.Email)
	user.Id, _ = result.LastInsertId()
}

// CreateUser creates a new user, saves it in the database, and returns it
func CreateUser(username, password, email string) *User {
	user := NewUser(username, password, email)
	user.Save()
	return user
}

// ListUsers retrieves and returns all users from the database
func ListUsers() Users {
	sql := "SELECT id, username, password, email FROM users"
	users := Users{}
	rows, _ := db.Query(sql)
	for rows.Next() {
		user := User{}
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
		users = append(users, user)
	}

	return users
}

// GetUser retrieves a single user by ID from the database
func GetUser(id int) *User {
	user := NewUser("", "", "")
	sql := "SELECT id, username, password, email FROM users WHERE id=?"
	rows, _ := db.Query(sql, id)
	for rows.Next() {
		rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	}

	return user
}

// update modifies an existing user in the database
func (user *User) update() {
	sql := "UPDATE users SET username=?, password=?, email=? WHERE id=?"
	db.Exec(sql, user.Username, user.Password, user.Email, user.Id)
}

// Save checks if a user already exists (by ID) and either inserts or updates it in the database
func (user *User) Save() {
	if user.Id == 0 {
		user.insert()
	} else {
		user.update()
	}
}

// Delete removes a user from the database by ID and prints a confirmation
func (user *User) Delete() {
	sql := "DELETE FROM users WHERE id=?"
	db.Exec(sql, user.Id)
	fmt.Printf("User with Id: %d deleted\n", user.Id)
}
