package main

import (
  "net/http"
  "html/template"
  "log"
)

func buildHandler() (w http.HandlerFunc) {
  mainHandler := func (w http.ResponseWriter, r *http.Request)  {
    // Home path
    if r.URL.Path == "/" {
    t, err := template.ParseFiles("html/story.html")
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
