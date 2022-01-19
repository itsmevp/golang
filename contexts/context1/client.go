package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	if resp, err := http.Get("http://localhost:9999"); err != nil {
		log.Fatal(err)
	} else {
		io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
	}
}
