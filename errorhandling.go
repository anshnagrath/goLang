package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Age       string
	Saying    []string
}

func main() {
	p := Person{
		FirstName: "ansh",
		LastName:  "nagrath",
		Saying:    []string{"shaken,not stirred", "want to be a go dev"},
	}
	bs, err := toJSON(p)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(bs))
	}

}
func toJSON(a interface{}) ([]byte, error) {
	fmt.Println(a)
	bs, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
		fmt.Println(bs)
		return []byte{}, err
	}
	fmt.Println(string(bs))
	return bs, nil

}
