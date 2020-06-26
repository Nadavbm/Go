package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is Home Page")
}
