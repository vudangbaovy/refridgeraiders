package main

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

func cookieStore() *sessions.CookieStore {
	cookieStore := sessions.NewCookieStore([]byte("secret-session-key"))
	return cookieStore
}

func login(w http.ResponseWriter, r *http.Request) {
	session, err := cookieStore().Get(r, "Cookie Name")
	if err != nil {
		log.Fatalln(err)
	}

	session.Values["authenticated"] = true
	// saves session to the cookie store
	err = session.Save(r, w)
	if err != nil {
		return
	}

}

func AuthenticatedStat(w http.ResponseWriter, r *http.Request) {
	session, err := cookieStore().Get(r, "Cookie Name")
	if err != nil {
		log.Fatalln(err)
	}

	auth := session.Values["authenticated"]
	if auth == true {
		w.WriteHeader(http.StatusOK)
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
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
}
