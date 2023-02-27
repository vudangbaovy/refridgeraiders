package main

/*
dataBase.go is the go file for the backend functions
use command "go run ." to run with main.go
*/

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// user profile definition
type UserProfile struct {
	gorm.Model
	Name       string
	Password   string
	AdminLevel uint8
	Allergies  string
}

// sets up Sqlite3 database
func connnectDB(dbName string) *gorm.DB {
	dbName += ".db"
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("conncected to test.db")
	}
	return db
}

// function wraps all of the auto migration calls: for future use
func buildTables(db *gorm.DB) {
	db.AutoMigrate(&UserProfile{})
}

func handleUserPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["title"]
	newUserProfile(username, password, allergies)
}

// adding users function: future use when more tables added
func addUser(addUser *UserProfile, db *gorm.DB) (bool, *UserProfile) {
	searchUser := UserProfile{}
	db.Where("Name = ?", addUser.Name).First(&searchUser)

	if searchUser.ID == 0 {
		db.Create(&addUser)
		return true, addUser
	}
	return false, &searchUser
}

func newUserProfile(inputUserName string, inputPassword string, inputAllergies string, db *gorm.DB) (bool, *UserProfile) {
	userAdded, account := addUser(&UserProfile{Name: inputUserName, Password: inputPassword, AdminLevel: 0, Allergies: inputAllergies}, db)
	return userAdded, account
}

func loginUser(inputUserName string, inputPassword string, db *gorm.DB) (bool, *UserProfile) {
	//function tests inputted username and password against database, returns true and user's profile struct if successful
	//returns false and empty struct if unsuccessful
	user := UserProfile{}
	db.Where("Name = ?", inputUserName).First(&user)
	if user.ID == 0 || user.Password != inputPassword {
		return false, &UserProfile{}
	}
	return true, &user
}

func testUserAdd(db *gorm.DB) {
	//adding users test code, works in command line
	numberOfEntries := 0
	fmt.Print("Enter Number of Entries: ")
	fmt.Scan(&numberOfEntries)

	insertedUsers := make([]UserProfile, numberOfEntries)
	for i := 0; i < numberOfEntries; i++ {
		fmt.Println("Enter User Data In the Following Format: UserName Password AdminLevel")
		fmt.Scan(&insertedUsers[i].Name, &insertedUsers[i].Password, &insertedUsers[i].AdminLevel)
		fmt.Println("Enter Any Allergies Seperated By a comma")
		fmt.Scan(&insertedUsers[i].Allergies)
		addUser(&insertedUsers[i], db)
	}
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

func testUserSearch(db *gorm.DB) {
	//Function is a framework for a future function to query database

	numberOfEntries := 0
	fmt.Print("Enter Number of Queries: ")
	fmt.Scan(&numberOfEntries)

	for i := 0; i < numberOfEntries; i++ {
		var searchName string
		var foundUser UserProfile
		fmt.Print("Enter UserName To Search: ")
		fmt.Scan(&searchName)
		db.Where("Name = ?", searchName).First(&foundUser)

		fmt.Println("\nStored Information:\nUserName: " + foundUser.Name + "\nPassword: " + foundUser.Password)
		fmt.Println("Allergies: " + foundUser.Allergies + "\nCreatedAt: " + foundUser.CreatedAt.String())
		fmt.Println()
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
