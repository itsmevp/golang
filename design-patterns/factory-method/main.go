package main

import (
	"fmt"
	"strings"
)

type person struct {
	name   string
	gender string
	age    int
}

// Factory method allows you to create
func newPerson(name string, gender string, age int) person {

	p := person{}

	if strings.ToLower(gender) == "male" {
		p.name = fmt.Sprintf("Mr. %s", name)
		p.age = age
	} else {
		p.name = fmt.Sprintf("Mr. %s", name)
		p.age = age
	}

	return p
}

func main() {
	p := newPerson("vishnu", "male", 30)
	fmt.Println(p)
}
