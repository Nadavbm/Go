package main

import (
	"encoding/json"
	"fmt"
)

type Album struct {
	Name string `json:"name"`
	Year int    `json:"year"`
	Band Band   `json:"band"`
}

type Band struct {
	Name  string `json:"name"`
	Genre string `json:"genre"`
}

func jsonMarshal() {
	band := Band{Name: "Mastodon", Genre: "Metal"}
	album := Album{Name: "Leviathan", Year: 2003, Band: band}
	byteArray, err := json.Marshal(album)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteArray))
}

func jsonUnmarshal() {
	jsonContent := `{"name":"Leviathan","year":2003,"band":{"name":"Mastodon","genre":"Metal"}}`
	var read Album
	err := json.Unmarshal([]byte(jsonContent), &read)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", read)
}

func main() {
	jsonMarshal()
	jsonUnmarshal()
}
