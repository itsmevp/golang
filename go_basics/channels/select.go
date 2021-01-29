package main

import (
	"fmt"
	"time"
)

func singapore(c chan string) {
	time.Sleep(1 * time.Second)
	c <- "hello from singapore!"
}

func ireland(c chan string) {
	time.Sleep(4 * time.Second)
	c <- "hello from ireland!"
}

func testDefaultCase(c chan string) {
	time.Sleep(10500 * time.Millisecond)
	c <- "hello from default case test"
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go singapore(c1)
	go ireland(c2)

	// The select statement blocks until one of the case statement is ready.

	select {
	case response1 := <-c1:
		fmt.Println(response1)
	case response2 := <-c2:
		fmt.Println(response2)
	}

	// The default case in a select statement is executed when none of the other case is ready.
	c3 := make(chan string)
	go testDefaultCase(c3)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case val := <-c3:
			fmt.Println("value received:", val)
			return
		default:
			fmt.Println("no value received")
		}
	}

	// When multiple cases in a select statement are ready, one of them will be executed at random.

	// Empty select
	// We know that the select statement will block until one of its cases is executed.
	// In this case the select statement doesn't have any cases and hence it will block forever resulting in a deadlock.
}
