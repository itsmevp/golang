package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

type Person struct {
	Name   string
	Age    int
	Pizzas []string
}

func hello() {
	fmt.Println("Hello World")
}

func main() {
	Pizzas := []string{"corn", "cheeze"}
	p1 := Person{
		Name:   "Vishnu",
		Age:    29,
		Pizzas: Pizzas,
	}

	// trim markers
	// {{-  remove all the trailing white spaces / horizontal tab / carriage return / newline
	// -}} remove all the leading white spaces / horizontal tab / carriage return / newline

	t1 := template.Must(template.New("hello.txt").Parse("Hello {{- .Name -}} your age is {{.Age }}\n"))
	// HelloVishnuyour age is 29
	t1.Execute(os.Stdout, p1)

	// register the functions to funcmap of type template.FuncMap
	funcMap := template.FuncMap{
		"title": strings.Title,
		"upper": strings.ToUpper,
		"lower": strings.ToLower,
	}
	t2 := template.Must(template.New("func.txt").Funcs(funcMap).Parse("{{title .}}\n{{upper .}}\n{{lower .}}\n"))
	t2.Execute(os.Stdout, "hello world")
}
