package main

import (
	"fmt"
	"log"
	"net/http"
)


func main(){
  // Generate Json
  var listofstories *Listofstories 
  listofstories = Generate_jsondata()

  for key, val := range *listofstories {
    fmt.Println("key is: ", key)
    fmt.Println("value is: ", val)
  }

  fmt.Println(*listofstories[])
  // fmt.Println(listofstories)

  // Starting an HTTP Server
  mainHandler := buildHandler()
  log.Fatal(http.ListenAndServe(":8080", mainHandler))
}

