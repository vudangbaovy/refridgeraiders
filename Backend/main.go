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

	fmt.Println("starting")
	host := "localhost:4200"
	if err := http.ListenAndServe(host, httpHandler()); err != nil {
		log.Fatalf("Failed to listen on %s: %v", host, err)
	}
	/*
		db := connnectDB("test")
		buildTables(db)
		testUserAdd(db)
		testUserSearch(db)
		testLoginUser(db)
		testSoftDelete(db)
		testReturnSoftDelete(db)
		testUpdate(db)
		testHardDelete(db)

		serverStart(db)
	*/
}
