package main

import "fmt"

/*
make vs new

new:
- returns a pointer
  - newly allocated
  - zeroed value
    - returns a pointer to a newly allocated zeroed value of type T

make:
- slices, channels, maps
- allocates memory
- initializes
  - puts empty string into values
- when it comes to make, you are making something, for example if you are making a slice,
you are actually making an underlying array
*/

func main() {
	var newSlice *[]int = new([]int)
	fmt.Println(newSlice) // nil

	makeSlice := make([]int, 5, 10)
	fmt.Println(makeSlice) // [0 0 0 0 0]
}
