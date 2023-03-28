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
	UserComments []UserComment `gorm:"foreignKey:UserRef"`
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

type UserCommentJson struct {
	Name       	string `json:"name"`
	Password   	string `json:"password"`
	RecipeName 	string `json:"recipeName"`
	Comment		string `json:"comment"`
}

type UserComment struct {
	gorm.Model
	UserRef		uint
	Name		string
	RecipeName 	string
	Comment		string
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
	db.AutoMigrate(UserProfile{}, UserComment{})
}

func UserRegisterPost(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var newUserJson UserLoginJson
	json.NewDecoder(r.Body).Decode(&newUserJson)
	user := UserProfile{Name: newUserJson.Name, Password: newUserJson.Password, Allergies: newUserJson.Allergies}
	addUser(&user, connectDB("test")) //test db name
	json.NewEncoder(w).Encode(newUserJson)
}

func UserPost(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var loginJson UserLoginJson
	json.NewDecoder(r.Body).Decode(&loginJson)

	valid, user := loginUser(loginJson.Name, loginJson.Password, connectDB("test")) //test db name
	if valid {
		loginJson = UserLoginJson{Name: user.Name, Password: user.Password, Allergies: user.Allergies}
		json.NewEncoder(w).Encode(&loginJson)
	} else {
		json.NewEncoder(w).Encode(&UserLoginJson{})
	}

}

func UserPut(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var loginJson UserLoginJson
	json.NewDecoder(r.Body).Decode(&loginJson)

	db := connectDB("test")//test db name
	valid, user := loginUser(loginJson.Name, loginJson.Password, db) //returns if valid user and the user profile
	if valid {
		loginJson = UserLoginJson{Name: user.Name, Password: user.Password, Allergies: user.Allergies}
		json.NewEncoder(w).Encode(&loginJson)
	} else {
		json.NewEncoder(w).Encode(&UserLoginJson{})
	}
}

func UserDelete(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var deleteJson UserLoginJson
	json.NewDecoder(r.Body).Decode(&deleteJson)

	db := connectDB("test")
	var deleteUser UserProfile
	db.Where("Name = ?", deleteJson.Name).First(&deleteUser)
	if deleteUser.ID != 0 && deleteJson.Password == deleteUser.Password {
		db.Delete(&deleteUser)
		json.NewEncoder(w).Encode(&deleteJson)
	} else {
		json.NewEncoder(w).Encode(&UserLoginJson{})
	}
}

// adding users function: future use when more tables added
func addUser(addUser *UserProfile, db *gorm.DB) (bool, *UserProfile) {
	searchUser := UserProfile{}
	err := db.Limit(1).Find("Name = ?", addUser.Name).First(&searchUser)

	if err.Error != nil {
		result := db.Omit("UserComment").Create(&addUser)
		fmt.Println("User Added  : ", addUser.Name, " : Rows effected : ", result.RowsAffected)
		return true, addUser
	}
	return false, &searchUser
}

func loginUser(inputUserName string, inputPassword string, db *gorm.DB) (bool, *UserProfile) {
	//function tests inputted username and password against database, returns true and user's profile struct if successful
	//returns false and empty struct if unsuccessful
	var user UserProfile

	fmt.Println("User Login  : Username:", inputUserName, " Password:", inputPassword)
	err := db.Where("Name = ?", inputUserName).First(&user)
	if err.Error != nil || user.Password != inputPassword {
		fmt.Println("Login Attempt Failed")
		return false, nil
	}
	return true, &user
}

func recipeComAddPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var commentJson UserCommentJson
	json.NewDecoder(r.Body).Decode(&commentJson)

	db := connectDB("test")
	valid, user := loginUser(commentJson.Name, commentJson.Password, db)
	if valid {
		db.Model(&user).Association("UserComments").Append(&UserComment{Name: commentJson.Name, RecipeName: commentJson.RecipeName, Comment: commentJson.Comment})
		count := db.Model(&user).Association("UserComments").Count()
		fmt.Println("Number of comments: ", count)
		json.NewEncoder(w).Encode(&commentJson)
	} else {
		json.NewEncoder(w).Encode(&UserCommentJson{})
	}
}

func recipeComPost (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var commentJson UserCommentJson
	json.NewDecoder(r.Body).Decode(&commentJson)

	db := connectDB("test")
	valid, user := loginUser(commentJson.Name, commentJson.Password, db)

	if valid {
		exists, com := recipeComHelper(commentJson.RecipeName, user, db)
		if exists {
			commentJson.Comment = com
		}
		json.NewEncoder(w).Encode(&commentJson)
	} else {
		json.NewEncoder(w).Encode(&UserCommentJson{})
	}
}

func recipeComHelper(targetRecipe string, user *UserProfile, db *gorm.DB)(bool, string){
	var comments []UserComment
	db.Model(&user).Association("UserComments").Find(&comments)
	for _, v := range comments {
		if v.RecipeName == targetRecipe {
			return true, v.Comment
		}
	}
	return false, ""
}
