package main

/*
dataBase.go is the go file for the backend functions
use command "go run ." to run with main.go
*/

import (
	"encoding/json"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// user profile definition
type UserProfile struct {
	gorm.Model
	Name       string
	Password   string
	Allergies  string
}

type UserProfileJson struct {
	Name       string `json:"name"`
	Password   string `json:"password"`
	Allergies  string `json:"allergies"`
}

// sets up Sqlite3 database
func connectDB(dbName string) *gorm.DB {
	dbName += ".db"
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		buildTables(db)
	return db
	}
}

// function wraps all of the auto migration calls: for future use
func buildTables(db *gorm.DB) {
	db.AutoMigrate(&UserProfile{})
}

func UserRegisterPost(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var newUserJson UserProfileJson
	json.NewDecoder(r.Body).Decode(&newUserJson)

	user := UserProfile{Name: newUserJson.Name, Password: newUserJson.Password, Allergies: newUserJson.Allergies}
	addUser(&user, connectDB("test")) //test db name
	json.NewEncoder(w).Encode(newUserJson)
}

func UserPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var loginJson UserProfileJson
	json.NewDecoder(r.Body).Decode(&loginJson)

	valid, user := loginUser(loginJson.Name, loginJson.Password, connectDB("test")) //test db name
	if valid {
		loginJson = UserProfileJson{Name: user.Name, Password: user.Password, Allergies: user.Allergies}
		json.NewEncoder(w).Encode(&loginJson)
	} else {
		json.NewEncoder(w).Encode(&UserProfileJson{})
	}

}

func UserPut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var loginJson UserProfileJson
	json.NewDecoder(r.Body).Decode(&loginJson)

	db := connectDB("test")//test db name
	valid, user := loginUser(loginJson.Name, loginJson.Password, db) //returns if valid user and the user profile
	if valid {
		loginJson = UserProfileJson{Name: user.Name, Password: user.Password, Allergies: user.Allergies}
		json.NewEncoder(w).Encode(&loginJson)
	} else {
		json.NewEncoder(w).Encode(&UserProfileJson{})
	}
}

func UserDelete(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var deleteJson UserProfileJson
	json.NewDecoder(r.Body).Decode(&deleteJson)

	db := connectDB("test")
	var deleteUser UserProfile
	db.Where("Name = ?", deleteJson.Name).First(&deleteUser)
	if deleteUser.ID != 0 && deleteJson.Password == deleteUser.Password {
		db.Delete(&deleteUser)
		json.NewEncoder(w).Encode(&deleteJson)
	} else {
		json.NewEncoder(w).Encode(&UserProfileJson{})
	}
}

// adding users function: future use when more tables added
func addUser(addUser *UserProfile, db *gorm.DB) (bool, *UserProfile) {
	searchUser := UserProfile{}
	err := db.Limit(1).Find("Name = ?", addUser.Name).First(&searchUser)

	if err.Error != nil {
		db.Create(&addUser)
		return true, addUser
	}
	return false, &searchUser
}

func loginUser(inputUserName string, inputPassword string, db *gorm.DB) (bool, *UserProfile) {
	//function tests inputted username and password against database, returns true and user's profile struct if successful
	//returns false and empty struct if unsuccessful
	user := UserProfile{}
	err := db.Limit(1).Find("Name = ?", inputUserName).First(&user)
	if err.Error == nil || user.Password != inputPassword {
		return false, &UserProfile{}
	}
	return true, &user
}
