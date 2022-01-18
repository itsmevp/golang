package main

import "fmt"

// An interface is a set of method signatures
// When a type provides definition for all the methods in the interface is said to implement the interface

type vowelsFinder interface {
	findVowels() []rune
}

type myString string

func (m myString) findVowels() []rune {
	var v []rune
	for _, r := range m {
		if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' {
			v = append(v, r)
		}
	}
	return v
}

// type assertion is used to get the underlying value of the interface
// i.(T)

func main() {
	var v vowelsFinder
	ms := myString("Hello World")
	v = ms
	fmt.Printf("%c\n", v.findVowels())

	var i interface{} = 10
	fmt.Println(i.(int))
	// fmt.Println(i.(string)) panic: interface conversion: interface {} is int, not string

	if _, ok := i.(string); !ok {
		fmt.Printf("%v is not a string\n", i)
	}

	// Type switching
	switch i.(type) {
	case string:
		fmt.Println("it's a string")
	case int:
		fmt.Println("it's an int")
	default:
		fmt.Println("unknown type")
	}
}
