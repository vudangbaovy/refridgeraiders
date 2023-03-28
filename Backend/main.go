package main

import (
	"fmt"
	"net/http"
)

/*
This is the code for the backend, the backend is split into api and database management
dataBase.go is the database side
use command "go run ." while in the backend directory for the go files to run
*/

func main() {

	db := connectDB("test")
	buildTables(db)
	
	if (true) {
		var user UserProfile
		connectDB("test").First(&user)
		fmt.Println(user.Name)
		fmt.Println(user.Password)

		var results [3]bool
		fmt.Println("\nRunning DB Tests...")
		results[0] = testUserAdd(connectDB("test"))
		results[1] = testUserSearch(connectDB("test"))

		host := "localhost:3000"
		go http.ListenAndServe(host, httpHandler())

		results[2] = testUserPost()
		fmt.Println("\nTest Results: ")
		for i, v := range results {
		fmt.Printf("Test %d %t\n", i, v)
		}
		return
	}

	host := "localhost:3000"
	http.ListenAndServe(host, httpHandler())

	
}
