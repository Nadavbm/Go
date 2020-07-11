### json package

javaScript object notation (JSON) is an open standard file format, and data interchange format, that uses human-readable text to store and transmit data objects consisting of attribute–value pairs and array data types (or any other serializable value). 

json can have four peimitive types (strings, numbers, booleans, and null) and two structured types (objects and arrays)

json structure contains names\values:

```{"name1": "value1", "name2":{name2-1:"value2-1"}, "name3": 3, "name4": [1,2,3,4], "name5": null } ```

json representation of a person:

```
{
  "name": "John Mclain",
  "dyingHarder": true,
  "age": 32,
  "address": {
    "streetAddress": "Ben Yehuda 11",
    "city": "Tel Aviv",
    "state": "IL",
    "postalCode": "10021"
  },
  "phoneNumbers": [
    {
      "type": "home",
      "number": "03-678126"
    },
    {
      "type": "office",
      "number": "03-12396768"
    }
  ],
}
```

json array:
```
[
    {
        "name": "Shimi Tavori",
        "age": 32
    }.
    {
        "name": "Margalit Sanani",
        "age": 27
    },
]
```

golang json package documentation - https://godoc.org/encoding/json

json to go - https://mholt.github.io/json-to-go/

##### marshal

```func Marshal(v interface{}) ([]byte, error)``` - Marshal returns the JSON encoding of v.

marshal structs to json array:

```
package main

import (
	"encoding/json"
	"fmt"
)

type Food struct {
	Name  string `json:"Name"`
	Price int    `json: "Price"`
}

func main() {
	food1 := Food{
		Name:  "Hummus",
		Price: 13,
	}
	food2 := Food{
		Name:  "Falafel",
		Price: 11,
	}
	food := []Food{food1, food2}

	menu, err := json.Marshal(food)
	if err != nil {
		fmt.Println("could not json marshal")
	}
	fmt.Println(string(menu))
}
```

##### unmarshal

```func Unmarshal(data []byte, v interface{}) error``` - Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v.

```
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

func main() {
	albumJSON := `{"name":"Leviathan","year":2003,"band":{"name":"Mastodon","genre":"Metal"}}`
	var album Album
	err := json.Unmarshal([]byte(albumJSON), &album)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", album)

}
```

##### encoder



##### decoder