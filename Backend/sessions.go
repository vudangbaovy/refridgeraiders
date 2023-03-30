package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret-session-key"))

func login(w http.ResponseWriter, r *http.Request) {

	// get username and password from input
	//password := "temp pass"
	//username := "temp user"

	// get the hashed password from the database using the username
	// hash:= db.
	hash := "temp variable"
	// authenticate userpassword with database hash
	err := compareHash("hello", hash)
	if err == true {
		// creates a new session or retrieves exisitng
		session, err2 := store.Get(r, "session-name")

		if err2 != nil {
			http.Error(w, err2.Error(), http.StatusInternalServerError)
			return
		}
		session.Values["authenticated"] = true
		// save session
		err3 := session.Save(r, w)
		if err3 != nil {
			http.Error(w, err3.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
	fmt.Println("Incorrect password")
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	// removes user authentication
	session.Values["authenticated"] = false
	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func checking(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden access", http.StatusForbidden)
		return
	}
}
