package main

import (
	"fmt"
	"gomysql/db"
	"gomysql/models"
)

func main() {
	db.Connect()
	//fmt.Println(db.ExistsTable("users"))
	//db.CreateTable(models.UserSchema, "users")
	//db.Ping()
	//db.TruncateTable("users")
	user := models.CreateUser("user1", "user123465", "user@email.com")
	// users := models.ListUsers()
	// fmt.Println(users)

	//user := models.GetUser(2)
	fmt.Println(user)
	// user.Username = "juan"
	// user.Password = "jun123"
	// user.Email = "juan@juan.com"
	// user.Save()
	//user.Delete()
	//db.TruncateTable("users")
	fmt.Println(models.ListUsers())
	db.Close()
}
