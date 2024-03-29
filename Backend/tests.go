package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"
	"gorm.io/gorm"
)

// test 0
func testDBAdd(db *gorm.DB) bool {
	//adding users test code
	numberOfEntries := uint(3)
	insertedUsers := make([]UserProfile, numberOfEntries)

	fmt.Println("\nTest 0 -------------------------------------")
	for i := uint(0); i < numberOfEntries; i++ { //starts at last id so no duplicates in db accidently get deleted
		insertedUsers[i].User = "UATest" + strconv.FormatUint(uint64(i), 10) //creates userprofiles and adds them to db
		insertedUsers[i].Password = "password" + strconv.FormatUint(uint64(i), 10)
		insertedUsers[i].Allergies = "Allergies" + strconv.FormatUint(uint64(i), 10)
		addUser(&insertedUsers[i], db)
	}

	for i := uint(0); i < numberOfEntries; i++ {
		var searchUser UserProfile
		err := db.Where("User = ?", "UATest"+strconv.FormatUint(uint64(i), 10)).First(&searchUser)
		//finds added users

		//tests that they have the same values
		if searchUser.User != insertedUsers[i].User || searchUser.Password != insertedUsers[i].Password || searchUser.Allergies != insertedUsers[i].Allergies || err.Error != nil {
			fmt.Println("Failed      : ", searchUser.User, " : Does Not Match ", insertedUsers[i].User)
			for _, v := range insertedUsers {
				result := db.Unscoped().Delete(&v) //deletes added users from db
				fmt.Println("user Deleted: ", v.User, " : Rows Affected : ", result.RowsAffected)
			}
			return false //failed
		}
	}

	for _, v := range insertedUsers {
		result := db.Unscoped().Delete(&v)
		fmt.Println("user Deleted: ", v.User, " : Rows Affected : ", result.RowsAffected)
	}
	return true //passed
}

// test 1
func testDBSearch(db *gorm.DB) bool {
	//testing searching users

	fmt.Println("\nTest 1 -------------------------------------")
	var insertedUser UserProfile
	insertedUser.User = "USTest1" //creates userprofiles and adds them to db
	insertedUser.Password = "password1"
	insertedUser.Allergies = "Allergies1"
	addUser(&insertedUser, db)

	var searchUser UserProfile
	err2 := db.Where("User = ?", "USTest1").First(&searchUser)
	//finds added user

	//tests that it has the same values
	if err2.Error != nil {
		result := db.Unscoped().Delete(&insertedUser) //deletes added users from db
		fmt.Println("user Deleted: ", insertedUser.User, " : Rows Affected : ", result.RowsAffected)
		return false //failed
	}

	result := db.Unscoped().Delete(&insertedUser)
	fmt.Println("user Deleted: ", insertedUser.User, " : Rows Affected : ", result.RowsAffected)
	return true //passed
}

// test 2 - looks up a pre existing user's allergies
func testAllergiesPost() bool {
	fmt.Println("\nTest 2 -------------------------------------")
	time.Sleep(100 * time.Millisecond)
	postBody, _ := json.Marshal(map[string]string{
		"user":      "Nick",
		"password":  "Pwe2",
		"allergies": "",
	})

	responseBody := bytes.NewBuffer(postBody)

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
	if string(body) == "{\"user\":\"Nick\",\"password\":\"Pwe2\",\"allergies\":\"Pie\"}\n" {
		return true
	}
	return false
}

// test 3 - looks up a pre existing note
func testNotesPOST() bool {
	fmt.Println("\nTest 3 -------------------------------------")

	postBody, _ := json.Marshal(map[string]string{
		"user":       "Nick",
		"password":   "Pwe2",
		"RecipeName": "Cake",
		"note":       "",
	})
	responseBody := bytes.NewBuffer(postBody)

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
	if string(body) == "{\"user\":\"Nick\",\"password\":\"Pwe2\",\"recipeName\":\"Cake\",\"note\":\"Too Much Sugar\"}\n" {
		return true
	}
	return false
}

// test 4 - looks up a pre existing user's first and last name
func testUserPOST() bool {
	fmt.Println("\nTest 4 -------------------------------------")

	client := &http.Client{}

	postBody, _ := json.Marshal(map[string]string{
		"user":     "Nick",
		"password": "Pwe2",
	})

	req, err := http.NewRequest(http.MethodPost, "http://localhost:3000/user", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Printf("Request Error: %s\n", err)
		return false
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Request Error: %s\n", err)
		return false
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Read Error: %s\n", err)
	}
	if string(body) == "{\"user\":\"Nick\",\"password\":\"Pwe2\",\"firstN\":\"Nicholas\",\"lastN\":\"Callahan\"}\n" {
		return true
	}
	return false
}

// test 5 - Changes a user's name
func testUserPUT() bool {
	fmt.Println("\nTest 5 -------------------------------------")

	client := &http.Client{}

	//first message
	postBody, _ := json.Marshal(map[string]string{
		"user":     "Nick",
		"password": "Pwe2",
		"firstN":   "George",
		"lastN":    "Washington",
	})

	//sends message
	req, err := http.NewRequest(http.MethodPut, "http://localhost:3000/user", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Printf("Request Error: %s\n", err)
		return false
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Request Error: %s\n", err)
		return false
	}

	//checks response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Read Error: %s\n", err)
	}
	if string(body) != "{\"user\":\"Nick\",\"password\":\"Pwe2\",\"firstN\":\"George\",\"lastN\":\"Washington\"}\n" {
		return false
	}

	//checks if db saved ealier change
	//second message
	postBody2, _ := json.Marshal(map[string]string{
		"user":     "Nick",
		"password": "Pwe2",
	})

	//sends second message
	req2, err := http.NewRequest(http.MethodPost, "http://localhost:3000/user", bytes.NewBuffer(postBody2))
	if err != nil {
		fmt.Printf("Request Error: %s\n", err)
		return false
	}
	req2.Header.Set("Content-Type", "application/json; charset=utf-8")
	res2, err := client.Do(req)
	if err != nil {
		fmt.Printf("Request Error: %s\n", err)
		return false
	}

	//checks response
	body2, err := ioutil.ReadAll(res2.Body)
	if err != nil {
		fmt.Printf("Read Error: %s\n", err)
	}
	if string(body2) == "{\"user\":\"Nick\",\"password\":\"Pwe2\",\"firstN\":\"George\",\"lastN\":\"Washington\"}\n" {
		return true
	} else {
		return false
	}
}

// startup Tests
func StartUpTest() {
	postBody, _ := json.Marshal(map[string]string{
		"user":      "Server",
		"password":  "Starting",
		"allergies": "Test",
	})
	responseBody := bytes.NewBuffer(postBody)

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
		if string(body) == "{\"user\":\"Server\",\"password\":\"Starting\",\"allergies\":\"Test\"}\n" {
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

func correctPassTest(t *testing.T) bool {
	// should return true as correct pass is given
	pass := "Password1"
	hash, err := hashedPass(pass)
	if err != nil {
		t.Fatalf("Generating hashed password failed", err)
	}

	checking := compareHash(pass, hash)
	fmt.Println("Passwords Match: ", checking)
	return checking
}

func incorrectPassTest(t *testing.T) bool {
	// this should return false as an incorrect pass is given
	pass := "Password1"
	hash, err := hashedPass(pass)
	if err != nil {
		t.Fatalf("Generating hashed password failed", err)
	}

	pass2 := "Password10"
	checking := compareHash(pass2, hash)
	fmt.Println("Passwords Match: ", checking)
	return !checking
}

func TestUserDelete() bool {
	fmt.Println("\nTest 10 -------------------------------------")

	client := &http.Client{}

	db := connectDB("test")
	var count1 int64
	var count2 int64
	db.Where("user = ?", "Nick").Find(&UserProfile{}).Count(&count1)

	//first message
	postBody, _ := json.Marshal(map[string]string{
		"user":     "Nick",
		"password": "Pwe2",
	})

	//sends message
	req, err := http.NewRequest(http.MethodDelete, "http://localhost:3000/user", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Printf("Request Error: %s\n", err)
		return false
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	_, err = client.Do(req)
	if err != nil {
		fmt.Printf("Request Error: %s\n", err)
		return false
	}

	//checks
	db.Where("user = ?", "Nick").Find(&UserProfile{}).Count(&count2)
	return count1 == (count2 + 1)
}

func testLogin() bool {

	// sending it as JSON
	postBody, _ := json.Marshal(map[string]string{
		"user":     "test123",
		"password": "hello",
	})
	responseBody := bytes.NewBuffer(postBody)

	http.Post("http://localhost:3000/user/register", "application/json", responseBody)
	responseBody = bytes.NewBuffer(postBody)

	res1, err := http.Post("http://localhost:3000/login", "application/json", responseBody)
	if err != nil {
		fmt.Printf("Request Error: %s\n", err)
		fmt.Println("Unable to login: 361")
		return false
	}

	defer res1.Body.Close()

	body, err := ioutil.ReadAll(res1.Body)
	if err != nil {
		fmt.Printf("Read Error: %s\n", err)
		fmt.Println("Unable to login: 370")
		return false
	} else {
		if strings.Contains(string(body), "test123") {
			fmt.Println("Successfully logged in")

			// delete the user
			client := &http.Client{}
			responseBody = bytes.NewBuffer(postBody)
			//sends message
			req, err := http.NewRequest(http.MethodDelete, "http://localhost:3000/user", responseBody)
			if err != nil {
				fmt.Printf("Request Error: %s\n", err)
				return false
			}
			req.Header.Set("Content-Type", "application/json; charset=utf-8")
			_, err = client.Do(req)
			if err != nil {
				fmt.Printf("Request Error: %s\n", err)
				return false
			}
			return true
		}
		fmt.Println("Failed To Login")
		return false
	}

}

func testBookmark()(bool) {
	fmt.Println("\nTest 11 -------------------------------------")
	client := &http.Client{}

	//first message
	postBody, _ := json.Marshal(map[string]string{
		"user":     "Nick",
		"password": "Pwe2",
		"bookmarks": "",
	})

	req, _ := http.NewRequest(http.MethodGet, "http://localhost:3000/logout", bytes.NewBuffer(postBody))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client.Do(req)

	//sends message
	req, err := http.NewRequest(http.MethodPost, "http://localhost:3000/bookmark", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Printf("Request Error: %s\n", err)
		return false
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Request Error: %s\n", err)
		return false
	}

	//checks response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Read Error: %s\n", err)
	}
	if !strings.Contains(string(body), "\"bookmarks\":\"\"") {
		fmt.Printf("string(body): %v\n", string(body))
		return false
	}
	



	//second message a put to add a bookmark
	postBody2, _ := json.Marshal(map[string]string{
		"user":     "Nick",
		"password": "Pwe2",
		"bookmarks": "KeyLimePie",
	})

	req, err = http.NewRequest(http.MethodPut, "http://localhost:3000/bookmark", bytes.NewBuffer(postBody2))
	if err != nil {
		fmt.Printf("Request Error: %s\n", err)
		return false
	}

	
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res, err = client.Do(req)
	if err != nil {
		fmt.Printf("Request Error: %s\n", err)
		return false
	}

	//checks response
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Read Error: %s\n", err)
	}
	if !strings.Contains(string(body), "KeyLimePie") {
		return false
	}





	postBody3, _ := json.Marshal(map[string]string{
		"user":     "Nick",
		"password": "Pwe2",
		"bookmarks": "Pie",
	})

	//third message adds another bookmark containing same name
	req, err = http.NewRequest(http.MethodPut, "http://localhost:3000/bookmark", bytes.NewBuffer(postBody3))
	if err != nil {
		fmt.Printf("Request Error: %s\n", err)
		return false
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res, err = client.Do(req)

	cookie := res.Cookies()

	if err != nil {
		fmt.Printf("Request Error: %s\n", err)
		return false
	}

	fmt.Print(res.Cookies())
	//checks response
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Read Error: %s\n", err)
	}
	if !strings.Contains(string(body), "Pie") {
		return false
	}


	postBody4, _ := json.Marshal(map[string]string{
		"user":     "",
		"password": "",
		"bookmarks": "",
	})

	//fourth message to check all
	req, err = http.NewRequest(http.MethodPost, "http://localhost:3000/bookmark", bytes.NewBuffer(postBody4))
	if err != nil {
		fmt.Printf("Request Error: %s\n", err)
		return false
	}
	
	for _, v := range cookie {
		if v.Value != "" {
			req.AddCookie(v)
		} else {
			fmt.Println(v.Value)
		}

	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res, err = client.Do(req)
	if err != nil {
		fmt.Printf("Request Error: %s\n", err)
		return false
	}

	//checks response
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Read Error: %s\n", err)
	}
	if !strings.Contains(string(body), "KeyLimePie,Pie") {
		return false
	}



	//fith message deletes pie
	req, err = http.NewRequest(http.MethodDelete, "http://localhost:3000/bookmark", bytes.NewBuffer(postBody3))
	if err != nil {
		fmt.Printf("Request Error: %s\n", err)
		return false
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res, err = client.Do(req)
	if err != nil {
		fmt.Printf("Request Error: %s\n", err)
		return false
	}

	//checks response
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Read Error: %s\n", err)
	}
	if !strings.Contains(string(body), "Pie") {
		return false
	}



	//sixth message checks to make sure we didnt delete wrong
	req, err = http.NewRequest(http.MethodPost, "http://localhost:3000/bookmark", bytes.NewBuffer(postBody3))
	if err != nil {
		fmt.Printf("Request Error: %s\n", err)
		return false
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res, err = client.Do(req)
	if err != nil {
		fmt.Printf("Request Error: %s\n", err)
		return false
	}

	//checks response
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Read Error: %s\n", err)
	}
	if !strings.Contains(string(body), "KeyLimePie") {
		return false
	}
	return true
}

func testLogout() bool {
	// sending it as JSON
	postBody, _ := json.Marshal(map[string]string{
		"user":     "test123",
		"password": "hello",
	})
	responseBody := bytes.NewBuffer(postBody)

	http.Post("http://localhost:3000/user/register", "application/json", responseBody)
	responseBody = bytes.NewBuffer(postBody)

	res1, err := http.Post("http://localhost:3000/login", "application/json", responseBody)
	if err != nil {
		fmt.Printf("Request Error: %s\n", err)
		fmt.Println("Unable to login: 361")
		return false
	}

	defer res1.Body.Close()

	body, err := ioutil.ReadAll(res1.Body)
	if err != nil {
		fmt.Printf("Read Error: %s\n", err)
		fmt.Println("Unable to login: 370")
		return false
	} else {
		if strings.Contains(string(body), "test123") {
			fmt.Println("Successfully logged in")
		} else {
			fmt.Println("Failed To Login")
		}
	}

	_, err1 := http.Get("http://localhost:3000/logout")
	if err1 != nil {
		fmt.Println("Unable to logout")
		return false
	}

	fmt.Println("Logout successful!")
	return true
}

// func correctPassDB(t *testing.T) bool {
// return
// }
func RunUnitTests(dbEmpty bool) {
	//bool parameter is for if the db is empty so a default entry can be added
	db := connectDB("test")
	buildTables(db)

	if dbEmpty {
		hashpass, _ := hashedPass("Pwe2")
		defaultUser := UserProfile{User: "Nick", Password: hashpass, FirstN: "Nicholas", LastN: "Callahan", Allergies: "Pie"}
		addUser(&defaultUser, db)
		db.Model(&defaultUser).Association("UserNotes").Append(&UserNote{User: defaultUser.User,
			RecipeName: "Cake", Note: "Too Much Sugar"})

	}

	var user UserProfile
	db.First(&user)
	//fmt.Println("Test Username: ", user.User)
	//fmt.Println("Test Password: ", user.Password)

	var results [12]bool

	fmt.Println("Running hash password tests")
	results[6] = correctPassTest(&testing.T{})
	results[7] = incorrectPassTest(&testing.T{})

	fmt.Println("\nRunning DB Tests...")
	results[0] = testDBAdd(db)
	results[1] = testDBSearch(db)

	//server tests
	host := "localhost:3000"
	go http.ListenAndServe(host, httpHandler())

	results[2] = testAllergiesPost()
	results[3] = testNotesPOST()
	results[5] = testUserPOST()
	results[4] = testUserPUT()
	fmt.Println("Running login tests")
	results[11] = testBookmark()
	results[8] = TestUserDelete()
	results[9] = testLogin()
	fmt.Println("Running logout tests")
	results[10] = testLogout()
	fmt.Println("\nTest Results: ")

	for i, v := range results {
		if v {
			fmt.Printf("Test %d Passed\n", i)
		} else {
			fmt.Printf("Test %d Failed\n", i)
		}
	}

}
