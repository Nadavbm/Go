package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
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
	dec := json.NewDecoder(strings.NewReader(jsonSTR))
	for {
		var albums []Album
		if err := dec.Decode(&albums); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		for i, v := range albums {
			fmt.Printf("number %d\n-------\nAlbum name: %s\nby the band - %s\nRecord year: %d\n---------------\n", i+1, v.Name, v.Band.Name, v.Year)
		}
	}
}
