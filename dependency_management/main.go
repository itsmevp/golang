package main

import "fmt"

type database interface {
	Create() string
	Update() string
}

// mysql implements database interface
type mysql struct{}

func (m mysql) Create() string {
	return "mysql create action invoked"
}

func (m mysql) Update() string {
	return "mysql update action invoked"
}

// postgress implements database interface
type postgress struct{}

func (m postgress) Create() string {
	return "postgress create action invoked"
}

func (m postgress) Update() string {
	return "postgress update action invoked"
}

// injecting dependencies to application  struct
type application struct {
	db database
}

func (a application) execute() {
	fmt.Println(a.db.Create())
	fmt.Println(a.db.Update())
}

func main() {
	a := application{
		db: postgress{},
	}

	/*
		a := application{
			db: mysql{},
		}
	*/

	// postgress create action invoked
	//postgress update action invoked
	a.execute()
}
