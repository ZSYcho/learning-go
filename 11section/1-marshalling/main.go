package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type user struct {
	Name     string `json:"name" xml:"name"`
	Age      int    `json:"age" xml:"age"`
	Phone    string `json:"phone" xml:"phone_number"`
	IsActive bool   `json:"active" xml:"active"`
}

func main() {

	var jane = user{
		Name:     "Jane",
		Age:      42,
		Phone:    "123-456-789",
		IsActive: true,
	}

	byteSlice, err := json.MarshalIndent(jane, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(byteSlice))
}
