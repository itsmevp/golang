package main

import (
	"fmt"
)

func main() {
	var c chan int
	if c == nil {
		fmt.Println(c==nil)
		fmt.Println("Channel is nil, going to define it")
		c = make(chan int)
		fmt.Printf("Channel c is of type: %T\n", c)
		fmt.Println(c==nil)
	}
}
