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

	//gets allergies from database
	router.HandleFunc("/allergies", AllergiesPost).Methods("POST")
		//json {user: "UsernameExample", password: "PasswordExample", allergies: ""} returns
		//{user: "UsernameExample", password: "PasswordExample", allergies: ",allergyExample1,allergyExample2"}

		//wrong input returns
		//{user: "", password: "", allergies: ""}

	//adds new allergy to list
	router.HandleFunc("/allergies", AllergiesPut).Methods("PUT")
		//json {user: "UsernameExample", password: "PasswordExample", allergies: "allergyExample3"} returns
		//{user: "UsernameExample", password: "PasswordExample", allergies: "allergyExample3"}

		//wrong input returns
		//{user: "", password: "", allergies: ""}

		//a post after this example put returns
		//{user: "UsernameExample", password: "PasswordExample", allergies: ",allergyExample1,allergyExample2,allergyExample3"}

	//deletes allergy from database
	router.HandleFunc("/allergies", AllergiesDelete).Methods("DELETE")
		//json {user: "UsernameExample", password: "PasswordExample", allergies: "allergyExample3"} returns
		//{user: "UsernameExample", password: "PasswordExample", allergies: "allergyExample3"}

		//wrong input returns
		//{user: "", password: "", allergies: ""}

		//a post after this example put returns
		//{user: "UsernameExample", password: "PasswordExample", allergies: ",allergyExample1,allergyExample2"}


	//user is added to database
	router.HandleFunc("/user/register", UserRegisterPost).Methods("POST")
		//json {user: "UsernameExample", password: "PasswordExample", firstN: "FirstNameExample", lastN: "LastNameExample"} returns
		//{user: "UsernameExample", password: "PasswordExample", firstN: "FirstNameExample", lastN: "LastNameExample"}

	//gets user first name and last name
	router.HandleFunc("/user", UserPOST).Methods("POST")
		//json {user: "UsernameExample", password: "PasswordExample", firstN: "", lastN: ""} returns
		//{user: "UsernameExample", password: "PasswordExample", firstN: "FirstNameExample", lastN: "LastNameExample"}

		//wrong input returns
		//{user: "", password: "", firstN: "", lastN: ""}

	//updates users first and last name
	router.HandleFunc("/user", UserPUT).Methods("PUT")
		//json {user: "UsernameExample", password: "PasswordExample", firstN: "FirstNameExample2", lastN: "LastNameExample2"} returns
		//{user: "UsernameExample", password: "PasswordExample", firstN: "FirstNameExample2", lastN: "LastNameExample2"}
		//
		//wrong input returns
		//{user: "", password: "", firstN: "", lastN: ""}
		//
		//a post after this example put returns
		//{user: "UsernameExample", password: "PasswordExample", firstN: "FirstNameExample2", lastN: "LastNameExample2"}

	//deletes user
	router.HandleFunc("/user", UserDelete).Methods("DELETE")
		//json {user: "UsernameExample", password: "PasswordExample", firstN: "", lastN: ""} returns
		//{user: "UsernameExample", password: "PasswordExample", firstN: "", lastN: ""}
		//
		//wrong input returns
		//{user: "", password: "", firstN: "", lastN: ""}
		//
		//a post after this example put returns
		//{user: "", password: "", firstN: "", lastN: ""}

	//creates note
	router.HandleFunc("/note/create", CreateNotePost).Methods("POST")
		//json {user: "UsernameExample", password: "PasswordExample", recipeName: "Pizza", note: "TOO MUCH CHEESE"} returns
		//json {user: "UsernameExample", password: "PasswordExample", recipeName: "Pizza", note: "TOO MUCH CHEESE"}
		//
		//wrong input returns
		//json {user: "", password: "", recipeName: "", note: ""}

	//retrieves note from database
	router.HandleFunc("/note", NotePost).Methods("POST")
		//json {user: "UsernameExample", password: "PasswordExample", recipeName: "Pizza", note: ""} returns
		//{user: "UsernameExample", password: "PasswordExample", recipeName: "Pizza", note: "TOO MUCH CHEESE"}
		//
		//wrong input returns
		//json {user: "", password: "", recipeName: "", note: ""}

	//updates a note
	router.HandleFunc("/note", NotePut).Methods("PUT")
		//json {user: "UsernameExample", password: "PasswordExample", recipeName: "Pizza", note: "Not enough Cheese"} returns
		//{user: "UsernameExample", password: "PasswordExample", recipeName: "Pizza", note: "Not enough Cheese"}
		//
		//wrong input returns
		//json {user: "", password: "", recipeName: "", note: ""}
		//
		//post after this example put returns
		//{user: "UsernameExample", password: "PasswordExample", recipeName: "Pizza", note: "Not enough Cheese"}

	//deletes note
	router.HandleFunc("/note", NoteDelete).Methods("DELETE")
		//json {user: "UsernameExample", password: "PasswordExample", recipeName: "Pizza", note: ""} returns
		//{user: "UsernameExample", password: "PasswordExample", recipeName: "Pizza", note: ""}
		//
		//wrong input returns
		//json {user: "", password: "", recipeName: "", note: ""}
		//
		//post after this example put returns
		//{user: "UsernameExample", password: "PasswordExample", recipeName: "Pizza", note: ""}

	//retrives bookmarks
	router.HandleFunc("/bookmark", BookmarkPost).Methods("POST")
		//json {user: "UsernameExample", password: "PasswordExample", bookmarks: ""} returns
		//{user: "UsernameExample", password: "PasswordExample", bookmarks: ",bookmarkExample"}
		//
		//wrong input returns
		//json {user: "", password: "", bookmarks: ""}

	//adds another bookmark
	router.HandleFunc("/bookmark", BookmarkPut).Methods("PUT")
		//json {user: "UsernameExample", password: "PasswordExample", bookmarks: "bookmarkExample2"} returns
		//{user: "UsernameExample", password: "PasswordExample", bookmarks: "bookmarkExample2"}
		//
		//wrong input returns
		//json {user: "", password: "", bookmarks: ""}
		//
		//post after this example put returns
		//{user: "UsernameExample", password: "PasswordExample", bookmarks: ",bookmarkExample,bookmarkExample2"}

	//deletes a bookmark
	router.HandleFunc("/bookmark", BookmarkDelete).Methods("DELETE")
		//json {user: "UsernameExample", password: "PasswordExample", bookmarks: "bookmarkExample2"} returns
		//{user: "UsernameExample", password: "PasswordExample", bookmarks: "bookmarkExample2"}
		//
		//wrong input returns
		//json {user: "", password: "", bookmarks: ""}
		//
		//post after this example put returns
		//{user: "UsernameExample", password: "PasswordExample", bookmarks: ",bookmarkExample"}

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
