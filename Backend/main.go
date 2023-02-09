package main

/*
This is the code for the backend, the backend is split into api and database management
dataBase.go is the database side
use command "go run ." while in the backend directory for the go files to run
*/

func main() {

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
}
