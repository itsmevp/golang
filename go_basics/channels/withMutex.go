package main

import (
	"fmt"
	"sync"
)

var x int

func increment(w *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	w.Done()
}

func main() {
	var (
		wg sync.WaitGroup
		mu sync.Mutex
	)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(&wg, &mu)
	}
	wg.Wait()
	fmt.Println(x)
}
