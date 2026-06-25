package main

import "fmt"

type Number interface {
	int | float64 | float32
}

// Sum is a generic function
func Sum[T Number](numbers ...T) T {
	var total T
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {

	grades := []int{90, 85}
	people := []string{"Alice", "Bob", "Charlie"}

	// len() is a generic function
	fmt.Println(len(grades), len(people))

	v := Sum(10, 20.96, 30.0)
	fmt.Printf("value is %v, type is %T\n", v, v)
}
