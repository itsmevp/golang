package main

import (
	"testing"
)

func TestGreet(t *testing.T) {
	result := greet("")
	if result != "Hello Dude!" {
		t.Fatalf("Expected %v but got %v\n","Hello Dude!",result)
	}
	t.Logf("Expected %v and got %v\n","Hello Dude!",result)
}
