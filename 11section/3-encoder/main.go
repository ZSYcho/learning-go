package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type user struct {
	Name     string `json:"name" xml:"name"`
	Age      int    `json:"age" xml:"age"`
	Phone    string `json:"phone" xml:"phone_number"`
	Password string `json:"-" xml:"-"`
	IsActive bool   `json:"active" xml:"active"`
}

func main() {

	u := user{
		Name:     "John Smith",
		Age:      42,
		IsActive: true,
	}

	// customize the destination you want to write like os.Stdout
	//enc := json.NewEncoder(os.Stdout)
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	if err := enc.Encode(&u); err != nil {
		log.Fatal(err)
	}

	// if you choose the Stdout, you don't need to handle the output, that's it
	fmt.Println(string(buf.Bytes()))
}
