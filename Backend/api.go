package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// httpHandler creates the backend HTTP router for queries, types,
// and serving the Angular frontend.
func httpHandler() http.Handler {
	router := mux.NewRouter()
	// Your REST API requests go here

	//allergies update or get allergy data
	router.HandleFunc("/allergies", AllergiesPost).Methods("POST")
	router.HandleFunc("/allergies", AllergiesPut).Methods("PUT")
	//user is user management in db
	router.HandleFunc("/user/register", UserRegisterPost).Methods("POST")
	router.HandleFunc("/user", UserDelete).Methods("DELETE")
	//note is personal comments on recipes
	router.HandleFunc("/note/create", CreateNotePost).Methods("POST")
	router.HandleFunc("/note", NotePost).Methods("POST")
	router.HandleFunc("/note", NotePut).Methods("PUT")

	//start up tests
	router.HandleFunc("/servertest", JsonTest).Methods("POST")

	// Add your routes here.
	// WARNING: this route must be the last route defined.

	router.PathPrefix("/").Handler(AngularHandler).Methods("GET")

	/**
	 * We need some headers to be statically prepended to every response.
	 */
	return handlers.LoggingHandler(os.Stdout,
		handlers.CORS(
			handlers.AllowCredentials(),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization",
				"DNT", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since",
				"Cache-Control", "Content-Range", "Range"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"http://localhost:4200"}),
			handlers.ExposedHeaders([]string{"DNT", "Keep-Alive", "User-Agent",
				"X-Requested-With", "If-Modified-Since", "Cache-Control",
				"Content-Type", "Content-Range", "Range", "Content-Disposition"}),
			handlers.MaxAge(86400),
		)(router))
}
