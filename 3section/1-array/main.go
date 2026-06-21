package main

import "fmt"

func main() {
	// arrays are very powerful, efficient but not flexible

	var numbers [5]int
	fmt.Printf("%+v\n", numbers)

	numbers[1] = 10
	numbers[2] = 20
	fmt.Printf("%+v\n", numbers)

	primes := [4]int{2, 3, 5, 7}
	fmt.Printf("%+v\n", primes)
	primes[3] = 11
	fmt.Printf("%+v\n", primes)

	for i := 0; i < len(primes); i++ {
		fmt.Printf("%d\n", primes[i])
	}

	var matrix [3][3]int // 2d arrays
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[0][2] = 3
	fmt.Printf("%+v\n", matrix)
}
