package main

import "fmt"

func main() {

	// Buffered channel means you can create a channel without receiver,
	// keep sending message until the channel is full

	// When you want to create a buffered channel, you need to specify capacity
	messages := make(chan string, 3)

	fmt.Println("Sending messages to buffered channel")
	messages <- "first message"
	messages <- "second message"
	messages <- "third message"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)

}
