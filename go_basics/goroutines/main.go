package main

import (
	"fmt"
	"sync"
)

/*
Placing the keyword "go" infront of the function will tell go to spin off, 
what's called as green thread and run that function in that green thread.
*/

/*
Most of the programming languages you probably heard of and worked with use OS threads.
What that means is they got individual function call stack dedicated to the execution of whatever the code handled by that thread.
OS threads are heavy, creation and destruction of OS thread is expensive.
*/

/*
Golang takes a slightly different approach in creating the threads.
Instead of creating these massive OS threads, golang creates a abstraction around these threads called as go routines.
Inside of the go runtime we got a scheduler that's gonna map these go routines on to these OS threads for periods of time.
With this abstraction the go routines can be started with very less stack space so that they can be re-allocated very quickly.
By default go is gonna give you the number of operating system threads equal to the number of cores availabled on the machine.
*/

/*
What a wait group does it helps to synchronize multiple go routines together.
*/

var wg = sync.WaitGroup{}

func main() {
	wg.Add(1)
	go sayHello(&wg)
	wg.Wait()
}

func sayHello(wg *sync.WaitGroup) {
	fmt.Println("Hello")
	wg.Done()
}
