package main

import (
	"fmt"
)

func simpleDefer() {
	fmt.Println("Function simpleDefer: Start")
	defer fmt.Println("Function simpleDefer: deferred")
	defer fmt.Println("Function simpleDefer: Middle")
	defer fmt.Println("Function simpleDefer: Middle")
	defer fmt.Println("Function simpleDefer: Middle")
}

func lifoSimpleDefer() {
	// LIFO like stack
	// means the last one pushed in will come out first
	fmt.Println("Function lifoSimpleDefer: Start")
	defer fmt.Println("First: deferred")
	defer fmt.Println("Second: deferred")
	defer fmt.Println("Function lifoSimpleDefer: Middle")
}
func main() {
	defer func() {
		fmt.Println("Before the return of main()")
	}()

	//// why we need the defer feature
	//// this is the process of handling the files
	//file, err := os.Create("./defer.text")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer file.Close()

	//simpleDefer()
	lifoSimpleDefer()

	fmt.Println("Last in main()")
}
