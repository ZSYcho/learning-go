package main

import "fmt"

func main() {

	var greeting string        // zero-value is an empty string ""
	greeting = "Hello, World!" // initialization

	fmt.Println(greeting)

	var count int
	count = 1
	fmt.Println(count)

	var isRunning bool
	isRunning = false
	fmt.Println(isRunning)

	var firstName, lastName string
	firstName = "Kevin"
	lastName = "Johnson"
	fmt.Println(firstName, lastName)

	email := "123@qq.com" // short declaration and initialization in the same time
	fmt.Println(email)

	age := 26
	fmt.Println(age)

	var year = 2026
	fmt.Println(year)
}
