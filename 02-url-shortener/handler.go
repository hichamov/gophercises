package shortener

import (
	"net/http"
	"log"
	"gopkg.in/yaml.v3"
)

type Route struct {
	Routes []map[string]string
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