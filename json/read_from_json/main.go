package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type People struct {
	Person []Person
}

type Person struct {
	Name    string    `json:"name"`
	Address []Address `json:"address"`
}

type Address struct {
	Street  string `json:"street"`
	Number  int    `json:"number"`
	City    string `json:"city"`
	Country string `json:"country"`
}

func main() {
	jsonFile, err := os.Open("file.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully open json file! \nThe people in the list are: \n")
	defer jsonFile.Close()

	bs, _ := ioutil.ReadAll(jsonFile)

	asString := string(bs)
	personArr, err := peopleUnmarshal(asString)
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	for i, v := range personArr {
		fmt.Println("Person #", i)
		fmt.Println("Name: ", v.Name, "\nLives at: ")
		for _, val := range v.Address {
			fmt.Println(val.City, val.Country, "At street", val.Street, "number", val.Number)
		}
	}

	// Print the whole json file
	var result []Person
	json.Unmarshal([]byte(asString), &result)
	fmt.Println("The full json file is: \n", result)
}

func peopleUnmarshal(jstr string) ([]Person, error) {
	bs := []byte(jstr)
	var people []Person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		return nil, err
	}
	return people, nil
}
