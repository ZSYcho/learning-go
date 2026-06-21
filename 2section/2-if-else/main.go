package main

import "fmt"

func main() {

	tmp := 25
	if tmp > 30 {
		fmt.Println("greater than 30")
	} else {
		fmt.Println("is less than 30")
	}

	score := 85
	if score > 85 {
		fmt.Println("Grade A")
	} else if score > 80 {
		fmt.Println("Grade B")
	} else if score > 70 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("failed")
	}

	userAccess := map[string]bool{
		"jane": true,
		"john": false,
	}

	if hasAccess, ok := userAccess["john"]; ok && hasAccess {
		fmt.Println("Jane can access the system")
	} else {
		fmt.Println("access not granted")
	}
}
