package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {

	wg.Add(2)

	ch := make(chan int, 5)

	go func(ch <-chan int) {
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 1
		ch <- 2
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()
}
