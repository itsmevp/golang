package main

import (
	"fmt"
)

func greet(name string) (result string) {
	if name == "" {
		return fmt.Sprint("Hello Dude!")
	}
	return fmt.Sprintf("Hello %v!",name)
}
