package main

import (
	"embed"
	"fmt"
	"log"
)

// enterprise application in Go
// ----------------------------
var name = "Johnson"

//go:embed public
var public embed.FS

func main() {

	//fs.ReadFile(public, "public/data.txt")
	data, err := public.ReadFile("public/data.txt") // put the right path
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))

	fmt.Println(name)
}
