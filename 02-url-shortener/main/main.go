package main

import (
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
	// myrandompackage "github.com/hichamov/gophercises/02-url-shortener"
)

type Route struct {
	Routes []map[string]string
}

func main() {

	// This is the fallback handler
	fallbackHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Home Page!"))
	})

	yamlfile, err := os.ReadFile("routes.yml")
	if err != nil {
		log.Println(err)
	}
	
	mainHandler, _ := YAMLHandler(yamlfile, fallbackHandler)
	
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

func YAMLHandler(yamlcontent []byte, fallback http.Handler) (http.HandlerFunc, error) {

	parsedyaml := parseYaml(yamlcontent)
	finalmap := buildMap(parsedyaml)

	return MapHandler(finalmap, fallback), nil
}

func parseYaml(yamlstring []byte) Route{
	var parsedyaml Route
	err := yaml.Unmarshal(yamlstring, &parsedyaml)
	if err != nil {
		log.Println(err)
	}
	return parsedyaml
}

func buildMap(r Route) map[string]string{
	finalmap := make(map[string]string)
	var gkey, gval string
	
	for _, val := range r.Routes {
		for mkey, mval := range val {

			if mkey == "name" {
				gkey = mval
			}

			if mkey == "target" {
				gval = mval
			}

			finalmap[gkey] = gval
		}
	}

	return finalmap
}