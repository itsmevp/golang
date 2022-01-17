package main

import "fmt"

// In UTF-8 encoding a code point can occupy more than one byte
// This is where rune saves us.
// rune is a builtin type in go and it has an alias of int32

func main() {
	a := "Hello World!"
	for k, v := range a {
		fmt.Printf("%d : %c : %v\n", k, v, v)
	}
}

// len(s) is used to find the number of bytes in the string and it doesn't return the string length.
// utf8.RuneCountInString(s) (n int)
// Strings are immutable
// To workaround this string immutability, strings are converted to a slice of runes.
// Then that slice is mutated with whatever changes are needed and converted back to a new string.

func mutate(s []rune) string {
	s[0] = 'a'
	return string(s)
}
