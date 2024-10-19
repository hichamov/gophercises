package main

import (
	_"fmt"
	"log"
	"net/http"
)


func main(){
  // Generate Json
  var listofstories Listofstories 
  listofstories = Generate_jsondata()

  // for key, val := range listofstories {
  //   fmt.Println("key is: ", key)
  //   fmt.Println("value is: ", val)
  // }

  // fmt.Println("The content of home story is:", listofstories["home"])

  // Starting an HTTP Server
  mainHandler := buildHandler(listofstories)
  log.Fatal(http.ListenAndServe(":8080", mainHandler))
}
