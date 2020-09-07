package main

import (
	"os"
	"fmt"
)

func main() {
	fileName := "/Users/v01/MyCodeRepos/github/golang/go_basics/file_handling/sample_file.txt"
	fileOpenFlags := os.O_RDWR | os.O_CREATE | os.O_APPEND
	f, err := os.OpenFile("/tmp/hello.json", fileOpenFlags, 0744)
	// f, err := os.Open(file_name)
	if err != nil {
		fmt.Printf("[Error] Failed to open %v \n", fileName)
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Printf("[Error] Failed to close %v \n", fileName)
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
