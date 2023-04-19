package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

type userLogin struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

func cookieStore() *sessions.CookieStore {
	cookieStore := sessions.NewCookieStore([]byte("secret-session-key"))
	return cookieStore
}

func login(w http.ResponseWriter, r *http.Request) {

	//username := r.PostFormValue("user")
	//pass := r.PostFormValue("password")

	var newUserJson userLogin
	json.NewDecoder(r.Body).Decode(&newUserJson)

	// authenticate the user
	checking, user := ValidateUser(newUserJson.User, newUserJson.Password, connectDB("test"))
	if !checking {
		fmt.Print("Incorrect credentials")
		return
	}

	session, err := cookieStore().Get(r, "Cookie Name")
	if err != nil {
		fmt.Print(err)
	}
	session.Values["user"] = newUserJson.User
	session.Values["authenticated"] = true
	// saves session to the cookie store
	err = session.Save(r, w)
	if err != nil {
		return
	}

	// return to frontend the user info
	//json.NewEncoder(w).Encode(user)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AuthenticatedStat(w http.ResponseWriter, r *http.Request) (bool, string) {
	session, err := cookieStore().Get(r, "Cookie Name")
	if err != nil {
		log.Fatalln(err)
	}

	auth := session.Values["authenticated"]
	user := session.Values["user"].(string)
	if auth == true {
		w.WriteHeader(http.StatusOK)
		return true, user
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return false, ""
	}

}

func logout(w http.ResponseWriter, r *http.Request) {
	session, err := cookieStore().Get(r, "Cookie Name")
	if err != nil {
		return
	}

	session.Values["user"] = ""
	session.Values["authenticated"] = false
	delete(session.Values, "user")
	err = session.Save(r, w)
	if err != nil {
		return
	}
	// Redirect the user to the login page or home page
	//http.Redirect(w, r, "/login", http.StatusSeeOther)
}
