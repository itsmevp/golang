package main

import (
	"fmt"
	"sync"
)

var x = 0

func add(wg *sync.WaitGroup) {

	x += 1

	wg.Done()
}

func main() {

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go add(&wg)
	}
	wg.Wait()
	fmt.Println(x)
}
