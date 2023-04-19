package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

func cookieStore() *sessions.CookieStore {
	cookieStore := sessions.NewCookieStore([]byte("secret-session-key"))
	return cookieStore
}

func login(w http.ResponseWriter, r *http.Request) {

	username := r.PostFormValue("user")
	pass := r.PostFormValue("password")

	// authenticate the user
	checking, user := ValidateUser(username, pass, connectDB("test"))
	if checking == false {
		log.Fatalf("Incorrect credentials")
		return
	}

	session, err := cookieStore().Get(r, "Cookie Name")
	if err != nil {
		log.Fatalln(err)
	}
	session.Values["user"] = username
	session.Values["authenticated"] = true
	// saves session to the cookie store
	err = session.Save(r, w)
	if err != nil {
		return
	}

	// return to frontend the user info
	json.NewEncoder(w).Encode(user)
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

	session.Values["authenticated"] = false
	err = session.Save(r, w)
	if err != nil {
		return
	}
	// Redirect the user to the login page or home page
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
