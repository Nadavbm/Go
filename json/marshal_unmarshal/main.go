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
	f1 := Food{
		Name:  "Hummus",
		Price: 13,
	}
	fmarsh := foodMarshal(&f1)
	fmt.Println(fmarsh)

	f2 := `[{"Name":"Hummus","Price":13},{"Name":"Falafel","Price":11}]`
	var err error
	funmarsh, err := foodUnmarshal(f2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(funmarsh)

	for i, v := range funmarsh {
		fmt.Println("Menu dish number", i+1)
		fmt.Println("Name:", v.Name, "\tCost:", v.Price)
	}
}

func foodMarshal(f *Food) string {
	food := f

	menu, err := json.Marshal(food)
	if err != nil {
		fmt.Println("could not json marshal")
	}
	return string(menu)
}

func foodUnmarshal(jstr string) ([]Food, error) {
	bs := []byte(jstr)
	var food []Food

	err := json.Unmarshal(bs, &food)
	if err != nil {
		return nil, err
	}
	return food, nil
}
