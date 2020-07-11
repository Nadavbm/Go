package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
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

func main() {
	albumsJSON := `[{"name":"Leviathan","year":2003,"band":{"name":"Mastodon","genre":"Stoner Metal"}},{"name":"Album of the Year","year":1994,"band":{"name":"Faith No More","genre":"Almost Metal"}}]`
	albums := []Album{
		Album{
			Name: "Leviathan",
			Year: 2003,
			Band: Band{
				Name:  "Mastodon",
				Genre: "Stoner Metal",
			},
		},
		{
			Name: "Album of the Year",
			Year: 1994,
			Band: Band{
				Name:  "Faith No More",
				Genre: "Almost Metal",
			},
		},
	}
	fmt.Println("Encoding ...")
	encodeAlbum(albums)
	fmt.Println("Decoding ...")
	decodeAlbum(albumsJSON)
}

func encodeAlbum(albums []Album) {
	err := json.NewEncoder(os.Stdout).Encode(albums)
	if err != nil {
		fmt.Println(err)
	}
}

func decodeAlbum(jsonSTR string) {
	err := json.NewDecoder(strings.NewReader(jsonSTR))
	if err != nil {
		fmt.Println(err)
	}
}
