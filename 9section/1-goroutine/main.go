package main

import (
	"fmt"
	"time"
)

func sayHello(message string, delay time.Duration) {
	time.Sleep(delay)
	fmt.Println("sayHello", message)
}

func main() {

	fmt.Println("hello world from Main() Goroutine")

	// this is a different process entirely w/ interacting with main()
	go sayHello("hello world 1", time.Second)
	go sayHello("hello world 2", time.Second)
	go sayHello("hello world from 2 seconds", 2*time.Second)
	go sayHello("hello world from 3 seconds", 3*time.Second)

	fmt.Println("Last message from Main() Goroutine")

	// We have to wait a while for another goroutine executing before the end of main() goroutine
	time.Sleep(2 * time.Second)

}
