package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

func buildHandler(listofStories Listofstories) (w http.HandlerFunc) {
  
  mainHandler := func (w http.ResponseWriter, r *http.Request)  {
    path := r.URL.Path
    fpath := strings.Trim(path, "/")

  if _, ok := listofStories[fpath]; ok {
    t, err := template.ParseFiles("html/story.html")
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    err = t.Execute(w,listofStories[fpath])
    if err != nil {
      log.Println(err)
    }
  }else {

    t, err := template.ParseFiles("html/index.html")
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    
    err = t.Execute(w, nil)
    if err != nil {
      log.Println(err)
    
    }

  }

}
  return mainHandler
}
