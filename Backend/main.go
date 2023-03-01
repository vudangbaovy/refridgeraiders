package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
This is the code for the backend, the backend is split into api and database management
dataBase.go is the database side
use command "go run ." while in the backend directory for the go files to run
*/

func main() {

	var results [3]bool
	fmt.Println("\nRunning DB Tests...")
	results[0] = testUserAdd(connectDB("test"))
	results[1] = testUserSearch(connectDB("test"))
	results[2] = testUserGet()
	fmt.Println("Finished DB Tests")
	for i, v := range results {
		fmt.Printf("Test %d %t\n", i, v)
	}

	var user UserProfile
	connectDB("test").First(&user)
	fmt.Println(user.Name)
	fmt.Println(user.Password)

	host := "localhost:4200"
	if err := http.ListenAndServe(host, httpHandler()); err != nil {
		log.Fatalf("Failed to listen on %s: %v", host, err)
	}
}
