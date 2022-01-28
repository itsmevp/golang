package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {

	wg.Add(2)

	c := make(chan int)

	// ch <-chan int receive only channel
	go func(ch <-chan int) {
		fmt.Println(<-c)
		wg.Done()
		// ch <- 23
		// invalid operation: cannot send to receive-only type <-chan
	}(c)
	// here we are passing a bi-directional channel to uni-directional channel
	// the runtime understands this and automatically casts to uni-directional channel

	// ch chan<- int send only channel
	go func(ch chan<- int) {
		c <- 34
		wg.Done()
		// <-ch
		// invalid operation: cannot receive from send-only channel ch
	}(c)

	wg.Wait()
}
