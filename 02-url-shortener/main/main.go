package main

import (
	"log"
	"net/http"
	// myrandompackage "github.com/hichamov/gophercises/02-url-shortener"
)

func main() {

	// This map will be used to create a handler
	urlmap := map[string]string{
		"/godoc": "https://go.dev/doc/",
		"/pydoc": "https://docs.python.org/3/",
	}

	// // This is the fallback handler
	fallbackHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Home Page!"))
	})

	mainHandler := MapHandler(urlmap, fallbackHandler)
	
	log.Fatal(http.ListenAndServe(":8080", mainHandler))
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusSeeOther)
		}

		fallback.ServeHTTP(w, r)
	}
}