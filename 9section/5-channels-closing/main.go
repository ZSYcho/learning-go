package main

import (
	"fmt"
	"sync"
)

func main() {

	jobs := make(chan int, 5)
	var wg sync.WaitGroup

	// Increment the counter outside goroutine (Rule 1)
	wg.Add(1)

	go func(wg *sync.WaitGroup) {
		// remember decrease the counter by calling wg.Done inside goroutine (Rule 2)
		defer wg.Done()
		for {
			r, ok := <-jobs
			if ok {
				fmt.Println("Got this message from jobs channel", r)
			} else {
				fmt.Println("Channel closed")
				return
			}
		}
	}(&wg)

	for i := 1; i <= 3; i++ {
		jobs <- i
		fmt.Println("Sending ", i)
	}

	close(jobs) // make the ok in `r, ok := <-jobs` be false

	wg.Wait() // Rule 3
}

func doubleChannelConcept() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			r, ok := <-jobs
			if ok {
				fmt.Println("Got this message from jobs channel", r)
			} else {
				fmt.Println("Channel closed")
				done <- true
				return
			}
		}
	}()

	for i := 1; i <= 3; i++ {
		jobs <- i
		fmt.Println("Sending ", i)
	}

	close(jobs) // make the ok in `r, ok := <-jobs` be false

	<-done
}
