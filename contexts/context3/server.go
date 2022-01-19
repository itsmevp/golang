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

/*

 */

func handler(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()
	log.Println("handler started")
	defer log.Println("handler ended")
	select {
	case <-time.After(5 * time.Second):
		w.Write([]byte("hello world!"))
	case <-r.Context().Done():
		log.Println(r.Context().Err().Error())
		http.Error(w, r.Context().Err().Error(), http.StatusInternalServerError)
	}
}
