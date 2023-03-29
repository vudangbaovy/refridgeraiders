package main

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret-session-key"))

func handler(w http.ResponseWriter, r *http.Request) {
	// creates a new session or retrieves exisitng
	session, err := store.Get(r, "session-name")

	// finish up authentication fns

	// if the user is authenticated
	session.Values["authenticated"] = true

	// save session
	err2 := session.Save(r, w)
	if err2 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
