package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

func config(numbers ...int) {
	if len(numbers) > 0 {
		first := numbers[0]
		fmt.Println("first number:", first)
	} else {
		fmt.Println("default number")
	}
}

func main() {

	fmt.Println(sum(1, 2, 3, 4))

	config()
}
