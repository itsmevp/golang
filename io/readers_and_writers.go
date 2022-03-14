// refer https://jesseduffield.com/Golang-IO-Cookbook/

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type MyString struct {
	reader io.Reader
}

func (m *MyString) Read(p []byte) (n int, err error) {
	n, err = m.reader.Read(p)
	if err != nil {
		return n, err
	}
	return
}

func main() {
	data := strings.NewReader("Hello World!")

	m := MyString{
		reader: data,
	}

	p := make([]byte, 2)

	for {
		n, err := m.Read(p)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(p[:n]))
	}
}
