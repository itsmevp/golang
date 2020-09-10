package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup) {
	x = x + 1
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg)
	}
	wg.Wait()
	fmt.Println(x)
}

/*
v01@Vishnus-Mac ~$ go run -race withOutMutex.go

==================
WARNING: DATA RACE
Read at 0x000001228380 by goroutine 8:
  main.increment()
      /Users/v01/withOutMutex.go:11 +0x3a

Previous write at 0x000001228380 by goroutine 7:
  main.increment()
      /Users/v01/withOutMutex.go:11 +0x56

Goroutine 8 (running) created at:
  main.main()
      /Users/v01/withOutMutex.go:19 +0xab

Goroutine 7 (finished) created at:
  main.main()
      /Users/v01/withOutMutex.go:19 +0xab
==================
1000
Found 1 data race(s)
exit status 66

~$ go run withOutMutex.go
907
*/
