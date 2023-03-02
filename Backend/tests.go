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

func testUserAdd(db *gorm.DB)(bool) {
	//adding users test code
	numberOfEntries := uint(3)
	insertedUsers := make([]UserProfile, numberOfEntries)

	topID := uint(0)
	err := db.Limit(1).Find("id = ?", 0)
	if err.Error == nil {
		var tempUser UserProfile
		db.Last(&tempUser)
		topID = tempUser.ID
	}

	for  i := uint(0); i < numberOfEntries; i++ { //starts at last id so no duplicates in db accidently get deleted
		insertedUsers[i].Name = "UATest" + strconv.FormatUint(uint64(i + topID + 1), 10)//creates userprofiles and adds them to db
		insertedUsers[i].Password = "password" + strconv.FormatUint(uint64(i + topID + 1), 10)
		insertedUsers[i].Allergies = "Allergies" + strconv.FormatUint(uint64(i + topID + 1), 10)
		addUser(&insertedUsers[i], db)
	}

	for i := uint(0); i < numberOfEntries; i++ {
		var searchUser UserProfile
		err := db.Where("id = ?", (i + topID + 1)).First(&searchUser)
		//finds added users 

		//tests that they have the same values
		if searchUser.Name != insertedUsers[i].Name || searchUser.Password != insertedUsers[i].Password || searchUser.Allergies != insertedUsers[i].Allergies || err.Error != nil {
			for _, v := range insertedUsers {
				db.Unscoped().Delete(&v)//deletes added users from db
			}
			return false//failed
		}
	}

	for _, v := range insertedUsers {
		db.Unscoped().Delete(&v)
	}
	return true//passed
}

func testUserSearch(db *gorm.DB)(bool) {
	//testing searching users

	
	topID := uint(0)
	err := db.Limit(1).Find("id = ?", 0)
	if err.Error == nil {
		var tempUser UserProfile
		db.Last(&tempUser)
		topID = tempUser.ID
	}


	var insertedUsers UserProfile
	insertedUsers.Name = "USTest" + strconv.FormatUint(uint64(1 + topID), 10)//creates userprofiles and adds them to db
	insertedUsers.Password = "password" + strconv.FormatUint(uint64(1 + topID), 10)
	insertedUsers.Allergies = "Allergies" + strconv.FormatUint(uint64(1 + topID), 10)
	addUser(&insertedUsers, db)

	var searchUser UserProfile
	err2 := db.Where("Name = ? AND id = ?", "USTest" + strconv.FormatUint(uint64(1 + topID), 10), 1 + topID).First(&searchUser)
	//finds added user

	//tests that it has the same values
	if   err2.Error != nil{
		db.Unscoped().Delete(&insertedUsers)//deletes added users from db
		return false//failed
	}

	db.Unscoped().Delete(&insertedUsers)
	return true//passed
}

func testUserGet()(bool) {

	time.Sleep(100 * time.Millisecond)
	postBody, _ := json.Marshal(map[string]string{
		"name": "Nick",
		"password": "Pwe2",
		"allergies": "Peanuts",
	}) 

	responseBody  := bytes.NewBuffer(postBody)

	res, err := http.Post("http://localhost:3000/Users/Register", "application/json", responseBody)
	if err != nil {
		fmt.Printf("Request Error: %s\n", err)
		return false
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Read Error: %s\n", err)
	 }
	fmt.Println(string(body))

	var testJson UserProfileJson
	if testJson.Name != "Nick" || testJson.Password != "Pwe2" || testJson.Allergies != "" {
		return false
	}

	var testUser UserProfile
	connectDB("test").Where("Name = ? AND Password = ?", "Nick", "Pwe2").First(&testUser)
	if testUser.Name != "Nick" || testUser.Password != "Pwe2" || testUser.Allergies != "" {
		return false
	}
	return true
}






func testLoginUser(db *gorm.DB) {
	//allows for command line tests of login function
	numberOfEntries := 0
	fmt.Print("Enter Number of Login Tests: ")
	fmt.Scan(&numberOfEntries)

	for i := 0; i < numberOfEntries; i++ {
		var loginName string
		var inputPassword string
		fmt.Print("Enter UserName To Login: ")
		fmt.Scan(&loginName)
		fmt.Print("Enter Password To Login: ")
		fmt.Scan(&inputPassword)
		validLogin, user := loginUser(loginName, inputPassword, db)

		if validLogin {
			fmt.Println("\nLogin Successful\nStored Information:\nUserName: " + user.Name + "\nPassword: " + user.Password)
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
		db.Where("Name = ?", dUserName).First(&deleteUser)
		db.Delete(&deleteUser)
		deleteUser = UserProfile{}
		db.Unscoped().Where("Name = ?", dUserName).First(&deleteUser)
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

		db.Unscoped().Where("Name = ?", rUserName).First(&returnedUser).Update("deleted_at", nil)

		returnedUser = UserProfile{}
		db.Where("Name = ?", rUserName).First(&returnedUser)

		fmt.Println("\n" + rUserName + " Has been returned to the database, printing stored data: ")
		fmt.Println("UserName: " + returnedUser.Name + "\nPassword: " + returnedUser.Password)
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

		db.Where("Name = ?", uUserName).First(&updateUser).Update(uFieldName, uValue)
		fmt.Println("\nStored Information:\nUserName: " + updateUser.Name + "\nPassword: " + updateUser.Password)
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

		db.Where("Name = ?", hdUserName).First(&hardDeleteUser)

		fmt.Println("\nUser Stored Information:\nUserName: " + hardDeleteUser.Name + "\nPassword: " + hardDeleteUser.Password)
		fmt.Println("Allergies: " + hardDeleteUser.Allergies + "\nCreatedAt: " + hardDeleteUser.CreatedAt.String() + "\n\nDeleting User")

		db.Unscoped().Delete(&hardDeleteUser)

		hardDeleteUser = UserProfile{}
		db.Where("Name = ?", hdUserName).First(&hardDeleteUser)

		if hardDeleteUser.ID == 0 {
			fmt.Println("User Was Successfully Deleted")
		} else {
			fmt.Println("User Was Unsuccessfully Deleted")
		}
	}
}