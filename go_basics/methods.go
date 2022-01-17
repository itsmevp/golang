package main

import "fmt"

// to define a method on a type, the definition of the receiver type and the definition
// of a method should present in the same package

// go is not a purley object oriented language and it does not support classes
// hence methods on types are the way to acheive similar behaviour
type Employee struct {
	FirstName string
	LastName  string
	Age       int
	Location
}

type Location struct {
	Country string
}

func (l Location) country() {
	fmt.Println(l.Country)
}

func (e Employee) changeName(name string) {
	e.FirstName = name
}

func (e *Employee) changeAge(age int) {
	e.Age = age
}

func main() {
	e1 := Employee{
		FirstName: "vishnu",
		LastName:  "pradeep",
		Age:       89,
		Location: Location{
			Country: "India",
		},
	}

	fmt.Println(e1.FirstName) // "vishnu"
	e1.changeName("VISHNU")   // "VISHNU"
	fmt.Println(e1.FirstName) // "vishnu"

	// pointer receivers can be used when changes made to the receiver,
	// inside the method should be visible to the caller.
	fmt.Println(e1.Age) // 89
	e1.changeAge(90)    // 90
	fmt.Println(e1.Age) // 90

	// Methods belonging to anonymous fields of a struct can be called as if they,
	// belong to the structure where the anonymous field is defined.
	e1.country()

	// A function with a value argument(s) will accept only value arguments
	// A method with a value receiver will accept both value receiver as well as pointer receiver
}
