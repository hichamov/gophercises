package main

import (
	"log"
	"net/http"
	"os"

	mycustomshortner "github.com/hichamov/gophercises/02-url-shortener"
)

func main() {

	// This is the fallback handler
	fallbackHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Home Page!"))
	})

	yamlfile, err := os.ReadFile("routes.yml")
	if err != nil {
		log.Println(err)
	}
	
	mainHandler, _ := mycustomshortner.YAMLHandler(yamlfile, fallbackHandler)
	
	log.Fatal(http.ListenAndServe(":8080", mainHandler))
}