package main

import "fmt"

// Employee struct is called named struct because it creates a new data type called Employee,
// using which Employee structs can be created
// If a strct type start with capital letter then it is an exported struct
// Similarly if a filed name starts with a capital letter then it can be accessed from other packages
type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

// strcts with anonymous fields
type Person struct {
	string
	int
}

type bar struct {
	location string
	pin      int
}

// promoted fields
type foo struct {
	name string
	age  int
	// fields that belong to an anonymous struct field in a struct are called promoted fields
	bar // here bar is an anonymous field, which is a struct
}

func main() {
	e1 := Employee{"Vishnu", "Pradeep", 24} // order of the fields matters
	fmt.Println(e1)

	e2 := Employee{
		FirstName: "vishnu",
		Age:       10,
		LastName:  "pradeep",
	} // order of fields doesn't matter
	fmt.Println(e2)

	// it is possible to create structs without creating data types
	// it is called anonymous structs
	e3 := struct {
		Name string
		Age  int
	}{
		Name: "vishnu",
		Age:  30,
	}
	fmt.Println(e3)

	//Even though anonymous fields do not have an explicit name,
	//by default the name of an anonymous field is the name of its type.
	p1 := Person{
		string: "vishnu",
		int:    23,
	}
	fmt.Println(p1.int)
	fmt.Println(p1.string)

	foo1 := foo{
		name: "vishnu",
		age:  29,
		bar: bar{
			location: "india",
			pin:      641113,
		},
	}
	// fields that belong to an anonymous struct field in a struct are called promoted fields
	fmt.Println(foo1.bar.location)
	fmt.Println(foo1.location)
	fmt.Println(foo1.pin)

	e11 := Employee{
		FirstName: "vishnu",
		LastName:  "s",
		Age:       29,
	}

	e22 := Employee{
		FirstName: "vishnu",
		LastName:  "s",
		Age:       29,
	}

	e33 := Employee{
		FirstName: "vishnu",
	}

	if e11 == e22 {
		fmt.Println("e11 and e22 are equal")
	} else {
		fmt.Println("e11 and e22 are not equal")
	}

	if e11 == e33 {
		fmt.Println("e11 and e33 are equal")
	} else {
		fmt.Println("e11 and e33 are not equal")
	}
}
