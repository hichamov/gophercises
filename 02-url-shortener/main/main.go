package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	myrandompackage "github.com/hichamov/gophercises/02-url-shortener"
)

type Scope struct {
	Project string
	Area    string
}

type Note struct {
	Title string
	Tags  []string
	Text  string
	Scope Scope
}

func main() {

	// Define a map of path and target URLs to be used

	routesmap := map[string]string{
		"/godoc": "https://go.dev/doc/",
		"/pythondoc": "https://docs.python.org/3/",
	}

	// The default route handler
	defaultHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	// // notes endpoint handler
	// http.HandleFunc("POST /notes", createNote)

	// // Test routing handler
	// http.HandleFunc("GET /routing", routeRequest)

	// Use the new handler
	finalHandler := http.HandleFunc("/hicham", myrandompackage.MapHandler(routesmap, defaultHandler))

	http.HandleFunc("/", finalHandler)
	// Ster server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// func routeRequest(w http.ResponseWriter, r *http.Request) {
// 	http.Redirect(w, r, "https://pkg.go.dev/", http.StatusSeeOther)
// }

// func createNote(w http.ResponseWriter, r *http.Request) {
// 	var note Note
// 	decoder := json.NewDecoder(r.Body)

// 	if err := decoder.Decode(&note); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	fmt.Fprintf(w, "Note: %+v\n", note)

// }
