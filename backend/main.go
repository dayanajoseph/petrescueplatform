package main

import (
	"fmt"
	"log"
	"net/http"
	"petrescueplatform/backend/api"   // Adjust this import path as needed
	"petrescueplatform/backend/model" // Adjust this import path as needed

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

const dsn = "root:Bella1211!@tcp(localhost:3306)/PetRescueDB"

func init() {
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	fmt.Println("Database connection successfully established")

	// Initialize the models package with the database connection
	model.Init(db)
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/user/create", api.CreateUser)
	http.HandleFunc("/user/login", api.LoginUser)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Pet Rescue Platform!")
}
