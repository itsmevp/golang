package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("127.0.0.1:9999", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("handler started")
	defer log.Println("handler ended")
	time.Sleep(5 * time.Second)
	w.Write([]byte("hello world!"))
}
