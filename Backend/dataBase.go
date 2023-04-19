package main

/*
dataBase.go is the go file for the backend functions
use command "go run ." to run with main.go
*/

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"log"
	
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// user profile definition
type UserProfile struct {
	gorm.Model
	UserNotes []UserNote `gorm:"foreignKey:UserRef"`
	UserBookMarks string
	User      string
	FirstN    string
	LastN     string
	Password  string
	Allergies string
}

// types are seperated into json to send only send specific data
type AllergiesJson struct {
	User      string `json:"user"`
	Password  string `json:"password"`
	Allergies string `json:"allergies"`
}

type LoginJson struct {
	User     string `json:"user"`
	Password string `json:"password"`
	FirstN   string `json:"firstN"`
	LastN    string `json:"lastN"`
}

type UserNoteJson struct {
	User       string `json:"user"`
	Password   string `json:"password"`
	RecipeName string `json:"recipeName"`
	Note       string `json:"note"`
}

type UserNote struct {
	gorm.Model
	UserRef    uint
	User       string
	RecipeName string
	Note       string
}

type BookMarkJson struct {
	User     string `json:"user"`
	Password string `json:"password"`
	UserBookMarks string `json:"bookmarks"`
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
	var newUserJson LoginJson
	json.NewDecoder(r.Body).Decode(&newUserJson)

	// change the password to be hashed
	hash, err := hashedPass(newUserJson.Password)
	if err != nil {
		fmt.Println("password unable to be hashed")
	}

	user := UserProfile{User: newUserJson.User, Password: hash,
		FirstN: newUserJson.FirstN, LastN: newUserJson.LastN, Allergies: "", UserBookMarks: ""}
	addUser(&user, connectDB("test"))
	//ValidateUserSessions(w, r, newUserJson.User, newUserJson.Password, connectDB("test"))
	json.NewEncoder(w).Encode(newUserJson)
}

func UserPOST(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var LUS LoginJson
	json.NewDecoder(r.Body).Decode(&LUS)

	db := connectDB("test")
	valid, user := ValidateUserSessions(w, r, LUS.User, LUS.Password, db)
	if valid {
		LUS.FirstN = user.FirstN
		LUS.LastN = user.LastN
		json.NewEncoder(w).Encode(&LUS)
	} else {
		w.WriteHeader(201)
	}
}

func UserPUT(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var LUS LoginJson
	json.NewDecoder(r.Body).Decode(&LUS)

	db := connectDB("test")
	valid, user := ValidateUserSessions(w, r, LUS.User, LUS.Password, db)
	if valid {
		db.Model(&user).Updates(UserProfile{FirstN: LUS.FirstN, LastN: LUS.LastN})
		json.NewEncoder(w).Encode(&LUS)
	} else {
		json.NewEncoder(w).Encode(&LoginJson{})
	}
}

func AllergiesPost(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var ARJ AllergiesJson
	json.NewDecoder(r.Body).Decode(&ARJ)

	valid, user := ValidateUserSessions(w, r, ARJ.User, ARJ.Password, connectDB("test"))

	if valid {
		ARJ = AllergiesJson{User: user.User, Password: ARJ.Password, Allergies: user.Allergies}
		json.NewEncoder(w).Encode(&ARJ)
	} else {
		json.NewEncoder(w).Encode(&AllergiesJson{})
	}
}

func AllergiesPut(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var AEJ AllergiesJson
	json.NewDecoder(r.Body).Decode(&AEJ)

	db := connectDB("test")
	valid, user := ValidateUserSessions(w, r, AEJ.User, AEJ.Password, db)
	if valid {
		exists, _ := StrHelper(user.Allergies, AEJ.Allergies)
		if !exists{
			newAllergies := user.Allergies + "," + AEJ.Allergies
			db.Model(&user).Update("Allergies", newAllergies)
			json.NewEncoder(w).Encode(&AEJ)
		}
	} else {
		json.NewEncoder(w).Encode(&AllergiesJson{})
	}
}

func AllergiesDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var AEJ AllergiesJson
	json.NewDecoder(r.Body).Decode(&AEJ)

	db := connectDB("test")
	valid, user := ValidateUserSessions(w, r, AEJ.User, AEJ.Password, db)

	if valid {
		exists, index := StrHelper(user.Allergies, AEJ.Allergies)
		if exists {
			newAllergies := user.Allergies[:index-1] + user.Allergies[index+len(AEJ.Allergies):]
			db.Model(&user).Update("Allergies", newAllergies)
			json.NewEncoder(w).Encode(&AEJ)
		} else {
			json.NewEncoder(w).Encode(&AllergiesJson{})
		}
	} else {
		json.NewEncoder(w).Encode(&AllergiesJson{})
	}
}

func UserDelete(w http.ResponseWriter, r *http.Request) {
	//deletes a user from the db
	w.Header().Set("Content-Type", "application/json")
	var deleteJson LoginJson
	json.NewDecoder(r.Body).Decode(&deleteJson)

	db := connectDB("test")
	valid, user := ValidateUserSessions(w, r, deleteJson.User, deleteJson.Password, db)

	if valid {
		db.Select("UserNotes").Delete(&UserProfile{}, user.ID)
		json.NewEncoder(w).Encode(&deleteJson)
	} else {
		json.NewEncoder(w).Encode(&AllergiesJson{})
	}
}

func addUser(addUser *UserProfile, db *gorm.DB) (bool, *UserProfile) {
	//adds users and returns bool and user profile type
	searchUser := UserProfile{}
	err := db.Limit(1).Find("User = ?", addUser.User).First(&searchUser)

	if err.Error != nil {
		result := db.Create(&addUser)
		fmt.Println("User Added  : ", addUser.User, " : Rows effected : ", result.RowsAffected)
		return true, addUser
	}
	return false, &UserProfile{}
}

func ValidateUser(inputUserName string, inputPassword string, db *gorm.DB) (bool, *UserProfile) {
	//function tests inputted username and password against database, returns true and user's profile struct if successful
	//returns false and empty struct if unsuccessful

	var user UserProfile

	//fmt.Println("User Login  : Username:", inputUserName, " Password:", inputPassword)
	err := db.Where("User = ?", inputUserName).First(&user)
	if err.Error != nil || !compareHash(inputPassword, user.Password) {
		fmt.Println("Login Attempt Failed")
		return false, nil
	}
	return true, &user
}

func ValidateUserSessions(w http.ResponseWriter, r *http.Request, inputUserName string, inputPassword string, db *gorm.DB) (bool, *UserProfile) {
	//function tests inputted username and password against database, returns true and user's profile struct if successful
	//returns false and empty struct if unsuccessful

	var user UserProfile

	session, err := cookieStore().Get(r, "Cookie-Name")
	if err != nil {
		log.Fatalln(err)
	}

	auth := session.Values["authenticated"]
	if auth == true {
		fmt.Println(session.Values["user"].(string))
		db.Where("User = ?", session.Values["user"].(string)).First(&user)
		return true, &user
	}

	
	//fmt.Println("User Login  : Username:", inputUserName, " Password:", inputPassword)
	err2 := db.Where("User = ?", inputUserName).First(&user)
	if err2.Error != nil || !compareHash(inputPassword, user.Password) {
		fmt.Println("Login Attempt Failed")
		session.Values["authenticated"] = false
		session.Values["user"] = ""
		session.Save(r, w)
		return false, nil
	}

	session.Values["authenticated"] = true
	session.Values["user"] = user.User
	session.Save(r, w)
	
	return true, &user
}

func CreateNotePost(w http.ResponseWriter, r *http.Request) {
	//function add a new userNote type to personal user profile, username, password, and recipe required in json body
	w.Header().Set("Content-Type", "application/json")
	var noteJson UserNoteJson
	json.NewDecoder(r.Body).Decode(&noteJson)

	db := connectDB("test")
	valid, user := ValidateUserSessions(w, r, noteJson.User, noteJson.Password, db)
	if valid {
		db.Model(&user).Association("UserNotes").Append(&UserNote{User: noteJson.User, RecipeName: noteJson.RecipeName, Note: noteJson.Note})
		count := db.Model(&user).Association("UserNotes").Count()
		fmt.Println("Number of notes: ", count)
		json.NewEncoder(w).Encode(&noteJson)
	} else {
		json.NewEncoder(w).Encode(&UserNoteJson{})
	}
}

func NotePost(w http.ResponseWriter, r *http.Request) {
	//function retrives a user's personal note on a recipe, requires a username, password and recipe as json
	w.Header().Set("Content-Type", "application/json")
	var noteJson UserNoteJson
	json.NewDecoder(r.Body).Decode(&noteJson)

	db := connectDB("test")
	valid, user := ValidateUserSessions(w, r, noteJson.User, noteJson.Password, db)

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

func NoteHelper(targetRecipe string, user *UserProfile, db *gorm.DB) (bool, string) {
	//helper function for recipeComPost, searches for targeted recipe in list of notes made by user
	var notes []UserNote
	db.Model(&user).Association("UserNotes").Find(&notes)
	for _, v := range notes {
		if v.RecipeName == targetRecipe {
			return true, v.Note
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
	valid, user := ValidateUserSessions(w, r, noteJson.User, noteJson.Password, db)

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

func NoteDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var noteJson UserNoteJson
	json.NewDecoder(r.Body).Decode(&noteJson)

	db := connectDB("test")
	valid, user := ValidateUserSessions(w, r, noteJson.User, noteJson.Password, db)

	if valid {
		var notes []UserNote
		db.Model(&user).Association("UserNotes").Find(&notes)
		for i, v := range notes {
			if v.RecipeName == noteJson.RecipeName {
				db.Model(&user).Association("UserNotes").Delete(notes[i])
				db.Delete(&notes[i])
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

func BookmarkPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var BMJson BookMarkJson
	json.NewDecoder(r.Body).Decode(&BMJson)

	db := connectDB("test")
	valid, user := ValidateUserSessions(w, r, BMJson.User, BMJson.Password, db)

	if valid {
		BMJson.UserBookMarks = user.UserBookMarks
		json.NewEncoder(w).Encode(&BMJson)
	} else {
		json.NewEncoder(w).Encode(&BookMarkJson{})
	}
}

func BookmarkPut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var BMJson BookMarkJson
	json.NewDecoder(r.Body).Decode(&BMJson)

	db := connectDB("test")
	valid, user := ValidateUserSessions(w, r, BMJson.User, BMJson.Password, db)

	if valid {
		exists, _ := StrHelper(user.UserBookMarks, BMJson.UserBookMarks)
		if ! exists{
			newBookMarks := user.UserBookMarks + "," + BMJson.UserBookMarks
			db.Model(&user).Update("UserBookMarks", newBookMarks)
			json.NewEncoder(w).Encode(&BMJson)
		}
	} else {
		json.NewEncoder(w).Encode(&BookMarkJson{})
	}
}

func BookmarkDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var BMJson BookMarkJson
	json.NewDecoder(r.Body).Decode(&BMJson)

	db := connectDB("test")
	valid, user := ValidateUserSessions(w, r, BMJson.User, BMJson.Password, db)

	if valid {
		exists, index := StrHelper(user.UserBookMarks, BMJson.UserBookMarks)
		if exists {
			newBookMarks := user.UserBookMarks[:index-1] + user.UserBookMarks[index+len(BMJson.UserBookMarks):]
			db.Model(&user).Update("UserBookMarks", newBookMarks)
			json.NewEncoder(w).Encode(&BMJson)
		} else {
			json.NewEncoder(w).Encode(&BookMarkJson{})
		}
	} else {
		json.NewEncoder(w).Encode(&BookMarkJson{})
	}
}

func StrHelper (uBM string, jBM string)(bool,int) {
	if !strings.Contains(uBM, jBM) {
		return false, 0
	} else {
		i := strings.Index(uBM, jBM)
		if uBM[i-1] != ',' || (len(uBM) != i+len(jBM)) && (uBM[i+len(jBM)] != ',') {
			exists, index := StrHelper(uBM[i+len(jBM):], jBM)
			return exists, index+i+len(jBM)-1
	}
		return true, i
}
}