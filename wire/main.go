package main

import "fmt"

type Message string

func NewMessage() Message {
	return Message("Hi there!")
}

type Greeter struct {
	Message Message
}

func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

func (g Greeter) greet() Message {
	return g.Message
}

type Event struct {
	Greeter Greeter
	Foo     Foo
}

func NewEvent(g Greeter, f Foo) Event {
	return Event{Greeter: g, Foo: f}
}

func (e Event) Start() {
	fmt.Println(e.Greeter.greet())
}

type Foo struct{}

func NewFoo() Foo {
	return Foo{}
}

func main() {
	//message := NewMessage()
	//greeter := NewGreeter(message)
	//event := NewEvent(greeter)
	//event.start()
	e := InitializeEvent()
	e.Start()
}
