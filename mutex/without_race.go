package main

import (
	"fmt"
	"sync"
)

var x = 0

func add(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x += 1
	m.Unlock()
	wg.Done()
}

func main() {

	wg := sync.WaitGroup{}
	m := sync.Mutex{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go add(&wg, &m)
	}
	wg.Wait()
	fmt.Println(x)
}
