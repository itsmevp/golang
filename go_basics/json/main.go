package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type student struct {
	Name    string
	Age     int
	Address string
}

type employee map[string]interface{}

func main() {
	s1 := student{
		Name:    "vishnu",
		Age:     10,
		Address: "coimbatore",
	}
	jsonData, err := json.Marshal(s1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(jsonData))

	e1 := employee{
		"Name":    "pradeep",
		"Age":     12,
		"Address": "coimbatore",
	}
	jsonData, err = json.Marshal(e1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(jsonData))

	type profile struct {
		UserName  string
		Followers int
		Prime     bool
	}

	type member struct {
		Fname, lname string
		Age          int
		Address      string
		profile
	}

	m1 := member{
		Fname:   "vishnu",
		lname:   "pradeep",
		Age:     10,
		Address: "coimbatore",
		profile: profile{
			UserName:  "v10",
			Followers: 1000,
			Prime:     false,
		},
	}
	jsonData, err = json.Marshal(m1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(jsonData))
}
