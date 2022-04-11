package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	input := make(chan int, 100)
	output := make(chan int, 100)

	for i := 0; i < 100; i++ {
		go worker(input, output)
	}

	for i := 0; i < 100; i++ {
		input <- i
	}

	for i := 0; i < 50; i++ {
		fmt.Println(<-output)
	}
	end := time.Since(start)
	fmt.Println("elapsed:", end)
}

func worker(input <-chan int, output chan<- int) {
	for n := range input {
		output <- addNumbers(n)
	}
}

func addNumbers(i int) int {
	time.Sleep(1 * time.Second)
	total := 0
	if i != 0 {
		total += i % 10
		i /= 10
	}
	return total + i
}
