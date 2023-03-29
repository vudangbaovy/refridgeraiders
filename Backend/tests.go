package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"encoding/json"
	"io/ioutil"
	"gorm.io/gorm"
)
//test 0
func testUserAdd(db *gorm.DB)(bool) {
	//adding users test code
	numberOfEntries := uint(3)
	insertedUsers := make([]UserProfile, numberOfEntries)

	fmt.Println("\nTest 0 -------------------------------------")
	for  i := uint(0); i < numberOfEntries; i++ { //starts at last id so no duplicates in db accidently get deleted
		insertedUsers[i].User = "UATest" + strconv.FormatUint(uint64(i), 10)//creates userprofiles and adds them to db
		insertedUsers[i].Password = "password" + strconv.FormatUint(uint64(i), 10)
		insertedUsers[i].Allergies = "Allergies" + strconv.FormatUint(uint64(i), 10)
		addUser(&insertedUsers[i], db)
	}

	for i := uint(0); i < numberOfEntries; i++ {
		var searchUser UserProfile
		err := db.Where("User = ?", "UATest" + strconv.FormatUint(uint64(i), 10)).First(&searchUser)
		//finds added users 

		//tests that they have the same values
		if searchUser.User != insertedUsers[i].User || searchUser.Password != insertedUsers[i].Password || searchUser.Allergies != insertedUsers[i].Allergies || err.Error != nil {
			fmt.Println("Failed      : ", searchUser.User, " : Does Not Match ", insertedUsers[i].User)
			for _, v := range insertedUsers {
				result := db.Unscoped().Delete(&v)//deletes added users from db
				fmt.Println("user Deleted: ", v.User, " : Rows Affected : ", result.RowsAffected)
			}
			return false//failed
		}
	}

	for _, v := range insertedUsers {
		result := db.Unscoped().Delete(&v)
		fmt.Println("user Deleted: ", v.User, " : Rows Affected : ", result.RowsAffected)
	}
	return true//passed
}
//test 1
func testUserSearch(db *gorm.DB)(bool) {
	//testing searching users

	fmt.Println("\nTest 1 -------------------------------------")
	var insertedUser UserProfile
	insertedUser.User = "USTest1"//creates userprofiles and adds them to db
	insertedUser.Password = "password1"
	insertedUser.Allergies = "Allergies1"
	addUser(&insertedUser, db)

	var searchUser UserProfile
	err2 := db.Where("User = ?", "USTest1").First(&searchUser)
	//finds added user

	//tests that it has the same values
	if   err2.Error != nil{
		result := db.Unscoped().Delete(&insertedUser)//deletes added users from db
		fmt.Println("user Deleted: ", insertedUser.User, " : Rows Affected : ", result.RowsAffected)
		return false//failed
	}

	result := db.Unscoped().Delete(&insertedUser)
	fmt.Println("user Deleted: ", insertedUser.User, " : Rows Affected : ", result.RowsAffected)
	return true//passed
}

//test 2
func testUserPost()(bool) {

	fmt.Println("\nTest 2 -------------------------------------")
	time.Sleep(100 * time.Millisecond)
	postBody, _ := json.Marshal(map[string]string{
		"user": "Nick",
		"password": "Pwe2",
		"allergies": "",
	}) 

	responseBody  := bytes.NewBuffer(postBody)

	res, err := http.Post("http://localhost:3000/allergies", "application/json", responseBody)
	if err != nil {
		fmt.Printf("Request Error: %s\n", err)
		return false
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Read Error: %s\n", err)
	}
	if string(body) == "{\"user\":\"Nick\",\"password\":\"Pwe2\",\"allergies\":\"Pie\"}\n"{
		return true
	}
	return false
}

//test 3
func testNotes()(bool) {
	fmt.Println("\nTest 3 -------------------------------------")

	postBody, _ := json.Marshal(map[string]string{
		"user": "Nick",
		"password": "Pwe2",
		"RecipeName": "Cake",
		"note": "",
	})
	responseBody  := bytes.NewBuffer(postBody)

	res, err := http.Post("http://localhost:3000/note", "application/json", responseBody)
	if err != nil {
		fmt.Printf("Request Error: %s\n", err)
		return false
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Read Error: %s\n", err)
	}
	if string(body) == "{\"user\":\"Nick\",\"password\":\"Pwe2\",\"recipeName\":\"Cake\",\"note\":\"Too Much Sugar\"}\n"{
		return true
	}
	return false
}

//startup Tests
func StartUpTest() {
	postBody, _ := json.Marshal(map[string]string{
		"user": "Server",
		"password": "Starting",
		"allergies": "Test",
	})
	responseBody  := bytes.NewBuffer(postBody)

	res, err := http.Post("http://localhost:3000/servertest", "application/json", responseBody)
	if err != nil {
		fmt.Printf("Request Error: %s\n", err)
		fmt.Println("Failed To Connect To server : 148")
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Read Error: %s\n", err)
		fmt.Println("Failed To Connect To server : 157")
		return
	} else {
		if string(body) == "{\"user\":\"Server\",\"password\":\"Starting\",\"allergies\":\"Test\"}\n"{
			fmt.Println("Connect To server")
			return
		}
		fmt.Println("Failed To Connect To server : 164")
		return
	}
}
func JsonTest(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		var JsonFrame AllergiesJson
		json.NewDecoder(r.Body).Decode(&JsonFrame)
	
		if JsonFrame.User == "Server" && JsonFrame.Password == "Starting" && JsonFrame.Allergies == "Test" {
			json.NewEncoder(w).Encode(&JsonFrame)
		} else {
			json.NewEncoder(w).Encode(&AllergiesJson{})
		}
}

func RunTests(dbEmpty bool) {

	db := connectDB("test")
	buildTables(db)

	if dbEmpty {
		defaultUser := UserProfile{User: "Nick", Password: "Pwe2", Allergies: "Pie"}
		addUser(&defaultUser,db)
		db.Model(&defaultUser).Association("UserNotes").Append(&UserNote{User: defaultUser.User, 
			RecipeName: "Cake", Note: "Too Much Sugar"})
		
	}

	var user UserProfile
	db.First(&user)
	fmt.Println("Test Username: ",user.User)
	fmt.Println("Test Password: ",user.Password)

	var results [4]bool
	fmt.Println("\nRunning DB Tests...")
	results[0] = testUserAdd(db)
	results[1] = testUserSearch(db)

	//server tests
	host := "localhost:3000"
	go http.ListenAndServe(host, httpHandler())

	results[2] = testUserPost()
	results[3] = testNotes()
	fmt.Println("\nTest Results: ")
	
	for i, v := range results {
		if v {
			fmt.Printf("Test %d Passed\n", i)
		} else {
			fmt.Printf("Test %d Failed\n", i)
		}
	}
}

//unused tests
func testLoginUser(db *gorm.DB) {
	//allows for command line tests of login function
	numberOfEntries := 0
	fmt.Print("Enter Number of Login Tests: ")
	fmt.Scan(&numberOfEntries)

	for i := 0; i < numberOfEntries; i++ {
		var loginName string
		var inputPassword string
		fmt.Print("Enter User To Login: ")
		fmt.Scan(&loginName)
		fmt.Print("Enter Password To Login: ")
		fmt.Scan(&inputPassword)
		validLogin, user := ValidateUser(loginName, inputPassword, db)

		if validLogin {
			fmt.Println("\nLogin Successful\nStored Information:\nUserName: " + user.User + "\nPassword: " + user.Password)
			fmt.Println("Allergies: " + user.Allergies + "\nCreatedAt: " + user.CreatedAt.String())
			fmt.Println()
		} else {
			fmt.Println("Invalid Login")
		}
	}
}

func testSoftDelete(db *gorm.DB) {
	//framework for future functions used to softDelete
	softDeletes := 0
	fmt.Print("Enter Number of softDeletes: ")
	fmt.Scan(&softDeletes)
	for i := 0; i < softDeletes; i++ {
		var deleteUser UserProfile
		var dUserName string
		fmt.Print("Enter userName To Be Soft Deleted: ")
		fmt.Scan(&dUserName)
		db.Where("User = ?", dUserName).First(&deleteUser)
		db.Delete(&deleteUser)
		deleteUser = UserProfile{}
		db.Unscoped().Where("User = ?", dUserName).First(&deleteUser)
		fmt.Println("\n" + dUserName + " was soft deleted at: " + deleteUser.DeletedAt.Time.String())
	}
}

func testReturnSoftDelete(db *gorm.DB) {
	//Demonstration/Framework function used to show how the database works with returning softdeleted data

	returns := 0
	fmt.Print("Enter Number of Returns: ")
	fmt.Scan(&returns)
	for i := 0; i < returns; i++ {
		var returnedUser UserProfile
		var rUserName string

		fmt.Print("Enter userName To Be Returned: ")
		fmt.Scan(&rUserName)

		db.Unscoped().Where("User = ?", rUserName).First(&returnedUser).Update("deleted_at", nil)

		returnedUser = UserProfile{}
		db.Where("User = ?", rUserName).First(&returnedUser)

		fmt.Println("\n" + rUserName + " Has been returned to the database, printing stored data: ")
		fmt.Println("User: " + returnedUser.User + "\nPassword: " + returnedUser.Password)
		fmt.Println("Allergies: " + returnedUser.Allergies + "\nCreatedAt: " + returnedUser.CreatedAt.String())
	}
}

func testUpdate(db *gorm.DB) {
	//function allows user to update specific fields from the commandline used to show how interaction with
	//database works/is a framework for future functions

	updates := 0
	fmt.Print("Enter Number of updates: ")
	fmt.Scan(&updates)
	for i := 0; i < updates; i++ {

		var uUserName string
		var uFieldName string
		var uValue string
		var updateUser UserProfile

		fmt.Print("\nEnter userName To Be Updated: ")
		fmt.Scan(&uUserName)
		fmt.Print("\nEnter Field To Be Updated: ")
		fmt.Scan(&uFieldName)
		fmt.Print("\nEnter Updated Value: ")
		fmt.Scan(&uValue)

		db.Where("User = ?", uUserName).First(&updateUser).Update(uFieldName, uValue)
		fmt.Println("\nStored Information:\nUserName: " + updateUser.User + "\nPassword: " + updateUser.Password)
		fmt.Println("Allergies: " + updateUser.Allergies + "\nCreatedAt: " + updateUser.CreatedAt.String())
		fmt.Println()
	}
}

func testHardDelete(db *gorm.DB) {
	//This function hard deletes values in the database
	//this can be used later as a framework to build future delete functions

	hardDeletes := 0
	fmt.Print("Enter Number of HardDeletes: ")
	fmt.Scan(&hardDeletes)

	for i := 0; i < hardDeletes; i++ {
		var hdUserName string
		var hardDeleteUser UserProfile
		fmt.Print("\nEnter userName of HardDelete: ")
		fmt.Scan(&hdUserName)

		db.Where("User = ?", hdUserName).First(&hardDeleteUser)

		fmt.Println("\nUser Stored Information:\nUserName: " + hardDeleteUser.User + "\nPassword: " + hardDeleteUser.Password)
		fmt.Println("Allergies: " + hardDeleteUser.Allergies + "\nCreatedAt: " + hardDeleteUser.CreatedAt.String() + "\n\nDeleting user")

		db.Unscoped().Delete(&hardDeleteUser)

		hardDeleteUser = UserProfile{}
		db.Where("User = ?", hdUserName).First(&hardDeleteUser)

		if hardDeleteUser.ID == 0 {
			fmt.Println("user Was Successfully Deleted")
		} else {
			fmt.Println("user Was Unsuccessfully Deleted")
		}
	}
}