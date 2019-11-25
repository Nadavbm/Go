package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type People struct {
	People []Person `json:"people"`
}

type Person struct {
	Name    string `json:"name"`
	Street  string `json:"street"`
	Number  int    `json:"number"`
	City    string `json:"city"`
	Country string `json:"country"`
}

func main() {
	jsonFile, err := os.Open("people.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully open json file! \nThe people in the list are: \n")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Print every field alone
	var people People
	json.Unmarshal(byteValue, &people)
	for i := 0; i < len(people.People); i++ {
		fmt.Println("Name: " + people.People[i].Name)
		fmt.Println("Street: " + people.People[i].Street)
		fmt.Println("Number: " + strconv.Itoa(people.People[i].Number))
		fmt.Println("City: " + people.People[i].City)
		fmt.Println("Country: " + people.People[i].Country + "\n")
	}

	// Print the whole json file
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	fmt.Println("The full json file is: \n", result["people"])
}
