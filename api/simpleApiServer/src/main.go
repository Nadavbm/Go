package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type Address struct {
	Street string `json:"Street"`
	City string `json:"City"`
	Country string `json:"Country"`
}

type Addresses []Address

func allAddresses(w http.ResponseWriter, r *http.Request) {
	addresses := Addresses{
		Address{Street:"Hakaktus", City:"Hadera", Country:"Israel"},
	}
	fmt.Println("Address list here")
	json.NewEncoder(w).Encode(addresses)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is home page mother fucker!!!")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/addresses", allAddresses)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
