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
	buildTables(connectDB("test"))
	if false {
		RunUnitTests(true)
	} else {
		host := "localhost:3000"
		fmt.Println("Server Starting...")
		go StartUpTest()
		http.ListenAndServe(host, httpHandler())
	}
}
