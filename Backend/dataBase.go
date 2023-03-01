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
	AdminLevel uint8
	Allergies  string
}

type UserProfileJson struct {
	Name       string `json:"name"`
	Password   string `json:"password"`
	AdminLevel uint8  `json:"adminLevel"`
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

func handleUserPost(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var newUserJson UserProfileJson
	json.NewDecoder(r.Body).Decode(&newUserJson)

	user := UserProfile{Name: newUserJson.Name, Password: newUserJson.Password, AdminLevel: newUserJson.AdminLevel, Allergies: newUserJson.Allergies}
	addUser(&user, connectDB("test")) //test db name
	json.NewEncoder(w).Encode(newUserJson)
}

func handleUserGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var loginJson UserProfileJson
	json.NewDecoder(r.Body).Decode(&loginJson)

	valid, user := loginUser(loginJson.Name, loginJson.Password, connectDB("test")) //test db name
	if valid {
		loginJson = UserProfileJson{Name: user.Name, Password: user.Password, AdminLevel: user.AdminLevel, Allergies: user.Allergies}
		json.NewEncoder(w).Encode(&loginJson)
	} else {
		json.NewEncoder(w).Encode(&UserProfileJson{})
	}

}

func handleUserPut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var loginJson UserProfileJson
	json.NewDecoder(r.Body).Decode(&loginJson)

	db := connectDB("test")
	var updateUser UserProfile
	valid, user := loginUser(loginJson.Name, loginJson.Password, db) //test db name
	if valid {
		db.Where("Name = ?", user.Name).First(&updateUser).Update("Allergies", loginJson.Allergies)
		db.Where("Name = ?", user.Name).First(&updateUser).Update("AdminLevel", loginJson.AdminLevel)
		loginJson = UserProfileJson{Name: updateUser.Name, Password: updateUser.Password, AdminLevel: updateUser.AdminLevel, Allergies: updateUser.Allergies}
		json.NewEncoder(w).Encode(&loginJson)
	} else {
		json.NewEncoder(w).Encode(&UserProfileJson{})
	}
}

func handleUserDelete(w http.ResponseWriter, r *http.Request) {

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
	err := db.Where("Name = ?", addUser.Name).First(&searchUser)

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
	db.Where("Name = ?", inputUserName).First(&user)
	if user.ID == 0 || user.Password != inputPassword {
		return false, &UserProfile{}
	}
	return true, &user
}
