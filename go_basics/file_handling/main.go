package main

import (
	"os"
	"fmt"
)

func main() {
	file_name := "/Users/v01/MyCodeRepos/github/golang/go_basics/file_handling/sample_file.txt"
	f, err := os.Open(file_name)
	if err != nil {
		fmt.Printf("[Error] Failed to open %v \n", file_name)
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Printf("[Error] Failed to close %v \n", file_name)
		}
	}()
	buff := make([]byte, 4)
	for {
		n, err := f.Read(buff)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(string(buff[0:n]))
	}
}
