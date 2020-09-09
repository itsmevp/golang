package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func calcCube(num int, c chan int) {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit * digit
		num /= 10
	}
	c <- sum
	wg.Done()
}

func calcSquare(num int, c chan int) {
	sum := 0
        for num != 0 {
                digit := num % 10
                sum += digit * digit
                num /= 10
        }
        c <- sum
        wg.Done()
}

func main() {
	sq := make(chan int)
	cu := make(chan int)
	go calcCube(123, cu)
	go calcSquare(1234, sq)
	wg.Add(2)
	fmt.Printf("The square is %v and the cube is %v\n", <-sq, <-cu)
	wg.Wait()
}
