package main

// These packages functions will be used later in the code

import (
	//        "fmt"
	"html/template"
	"log"
	"net/http"
	//        "path/filepath"
	//	"os"
)

type Person struct {
	First string
	Last  string
}

var (
	templates = template.Must(template.ParseFiles("static/index.html"))
)

func main() {
	// All html files will be added to static, the following is file server function for html templates
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// The handlers here will be used to route between the web pages
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)

	// This will output only in server side
	log.Println("Listening...")
	// This is the listen and serve code block as a condition in case of an error
	// the Log fatal will so errors in case the server cannot listen or server pages correctly
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Define the template
	tmpl := template.Must(template.ParseFiles("static/index.html"))

	// If client post first name last name return values
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	// Definition of the form and the struct
	details := Person{
		First: r.FormValue("first"),
		Last:  r.FormValue("last"),
	}

	_ = details

	// In html there is a case for Success in submiting the form and then return Hello World {{.First}} name of the person
	tmpl.Execute(w, struct{ Success bool }{true})
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("about.html")
	t.Execute(w, r)
}
