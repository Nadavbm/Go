package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/gorilla/mux"
)

type Text struct {
	Id		string 		`json:"id"`
	Text	string		`json:"text"`
}

var Texts []Text

func main() {
	Texts = []Text{
		Text{Id: "1", Text: "Welcome to home page!" },
		}
	handleRequests()
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	router.HandleFunc("/text", addText).Methods("POST")
	
	log.Fatal(http.ListenAndServe(":8081", router))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to home page")
	json.NewEncoder(w).Encode(Texts)
}

func addText(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var text Text
	json.Unmarshal(reqBody, &text)
	Texts = append(Texts, text)
	json.NewEncoder(w).Encode(text)
}
