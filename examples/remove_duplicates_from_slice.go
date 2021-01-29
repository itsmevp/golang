package main

import (
	"fmt"
)

func remove_duplicates(slc []int) []int {
	m := make(map[int]bool)
	s := []int{}
	for _, v := range slc {
		_, ok := m[v]
		if !ok {
			m[v] = true
			s = append(s, v)
		}
	}
	return s
}

func main() {
	x := []int{1, 2, 3, 3, 5, 8, 6, 2}
	fmt.Println(remove_duplicates(x))	//[1 2 3 5 8 6]
}
