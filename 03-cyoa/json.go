package main

import (
  "flag"
  "os"
  "log"
  "encoding/json"
)

type Listofstories map[string]Story

type Story struct {
	Title   string   `json:"title"`  
	Story   []string `json:"story"`  
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"` 
}

func Generate_jsondata() (*Listofstories){
    
  filename := flag.String("filename", "jsondata/gopher.json","This Specifies the json filename")
  flag.Parse()

  var stoireslist Listofstories
  storycontent, err := os.ReadFile(*filename)
  err = json.Unmarshal(storycontent, &stoireslist)

	if err != nil {
		log.Println(err)
	}
  
  return &stoireslist
}
