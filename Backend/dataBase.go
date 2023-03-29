package main

/*
dataBase.go is the go file for the backend functions
use command "go run ." to run with main.go
*/

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// user profile definition
type UserProfile struct {
	gorm.Model
	UserNotes []UserNote `gorm:"foreignKey:UserRef"`
	Name       	string
	Password   	string
	Allergies  	string
}
//types are seperated into json to send only send specific data
type UserLoginJson struct {
	Name       	string `json:"name"`
	Password   	string `json:"password"`
	Allergies  	string `json:"allergies"`
}

type UserNoteJson struct {
	Name       	string `json:"name"`
	Password   	string `json:"password"`
	RecipeName 	string `json:"recipeName"`
	Note		string `json:"note"`
}

type UserNote struct {
	gorm.Model
	UserRef		uint
	Name		string
	RecipeName 	string
	Note		string
}

// sets up Sqlite3 database
func connectDB(dbName string) *gorm.DB {
	dbName += ".db"
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

// function wraps all of the auto migration calls: for future use
func buildTables(db *gorm.DB) {
	db.AutoMigrate(UserProfile{}, UserNote{})
}

func UserRegisterPost(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var newUserJson UserLoginJson
	json.NewDecoder(r.Body).Decode(&newUserJson)
	user := UserProfile{Name: newUserJson.Name, Password: newUserJson.Password, Allergies: newUserJson.Allergies}
	addUser(&user, connectDB("test"))
	json.NewEncoder(w).Encode(newUserJson)
}

func AllergiesPost(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var loginJson UserLoginJson
	json.NewDecoder(r.Body).Decode(&loginJson)

	valid, user := ValidateUser(loginJson.Name, loginJson.Password, connectDB("test"))
	if valid {
		loginJson = UserLoginJson{Name: user.Name, Password: user.Password, Allergies: user.Allergies}
		json.NewEncoder(w).Encode(&loginJson)
	} else {
		json.NewEncoder(w).Encode(&UserLoginJson{})
	}
}

func AllergiesPut(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var loginJson UserLoginJson
	json.NewDecoder(r.Body).Decode(&loginJson)

	db := connectDB("test")
	valid, user := ValidateUser(loginJson.Name, loginJson.Password, db)
	if valid {
		loginJson = UserLoginJson{Name: user.Name, Password: user.Password, Allergies: user.Allergies}
		json.NewEncoder(w).Encode(&loginJson)
	} else {
		json.NewEncoder(w).Encode(&UserLoginJson{})
	}
}

func UserDelete(w http.ResponseWriter, r *http.Request) {
	//deletes a user from the db
	w.Header().Set("Content-Type", "application/json")
	var deleteJson UserLoginJson
	json.NewDecoder(r.Body).Decode(&deleteJson)

	db := connectDB("test")
	valid, user := ValidateUser(deleteJson.Name, deleteJson.Password, db)

	if valid {
		db.Unscoped().Delete(&user)
		json.NewEncoder(w).Encode(&user)
	} else {
		json.NewEncoder(w).Encode(&UserProfile{})
	}
}

func addUser(addUser *UserProfile, db *gorm.DB) (bool, *UserProfile) {
	//adds users and returns bool and user profile type
	searchUser := UserProfile{}
	err := db.Limit(1).Find("Name = ?", addUser.Name).First(&searchUser)

	if err.Error != nil {
		result := db.Omit("UserNote").Create(&addUser)
		fmt.Println("User Added  : ", addUser.Name, " : Rows effected : ", result.RowsAffected)
		return true, addUser
	}
	return false, &UserProfile{}
}

func ValidateUser(inputUserName string, inputPassword string, db *gorm.DB) (bool, *UserProfile) {
	//function tests inputted username and password against database, returns true and user's profile struct if successful
	//returns false and empty struct if unsuccessful
	var user UserProfile

	//fmt.Println("User Login  : Username:", inputUserName, " Password:", inputPassword)
	err := db.Where("Name = ?", inputUserName).First(&user)
	if err.Error != nil || user.Password != inputPassword {
		fmt.Println("Login Attempt Failed")
		return false, nil
	}
	return true, &user
}

func CreateNotePost(w http.ResponseWriter, r *http.Request) {
	//function add a new userNote type to personal user profile, username, password, and recipe required in json body
	w.Header().Set("Content-Type", "application/json")
	var noteJson UserNoteJson
	json.NewDecoder(r.Body).Decode(&noteJson)

	db := connectDB("test")
	valid, user := ValidateUser(noteJson.Name, noteJson.Password, db)
	if valid {
		db.Model(&user).Association("UserNotes").Append(&UserNote{Name: noteJson.Name, RecipeName: noteJson.RecipeName, Note: noteJson.Note})
		count := db.Model(&user).Association("UserNotes").Count()
		fmt.Println("Number of notes: ", count)
		json.NewEncoder(w).Encode(&noteJson)
	} else {
		json.NewEncoder(w).Encode(&UserNoteJson{})
	}
}

func NotePost (w http.ResponseWriter, r *http.Request) {
	//function retrives a user's personal note on a recipe, requires a username, password and recipe as json
	w.Header().Set("Content-Type", "application/json")
	var noteJson UserNoteJson
	json.NewDecoder(r.Body).Decode(&noteJson)

	db := connectDB("test")
	valid, user := ValidateUser(noteJson.Name, noteJson.Password, db)

	if valid {
		exists, note := NoteHelper(noteJson.RecipeName, user, db)
		if exists {
			noteJson.Note = note
		}
		json.NewEncoder(w).Encode(&noteJson)
	} else {
		json.NewEncoder(w).Encode(&UserNoteJson{})
	}
}

func NoteHelper(targetRecipe string, user *UserProfile, db *gorm.DB)(bool, string){
	//helper function for recipeComPost, searches for targeted recipe in list of notes made by user
	var notes []UserNote
	db.Model(&user).Association("UserNotes").Find(&notes)
	for _, v := range notes {
		if v.RecipeName == targetRecipe {
			return true,  v.Note
		}
	}
	return false, ""
}

func NotePut(w http.ResponseWriter, r *http.Request) {
	//function takes json object with updated note value
	w.Header().Set("Content-Type", "application/json")
	var noteJson UserNoteJson
	json.NewDecoder(r.Body).Decode(&noteJson)

	db := connectDB("test")
	valid, user := ValidateUser(noteJson.Name, noteJson.Password, db)
	
	if valid {
		var notes []UserNote
		db.Model(&user).Association("UserNotes").Find(&notes)
		for i, v := range notes {
			if v.RecipeName == noteJson.RecipeName {
				notes[i].Note = noteJson.Note
				db.Model(&user).Association("UserNotes").Replace(notes)
				db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user)
				json.NewEncoder(w).Encode(&noteJson)
				return
			}
		}
		json.NewEncoder(w).Encode(&UserNoteJson{})
	} else {
		json.NewEncoder(w).Encode(&UserNoteJson{})
	}
}
