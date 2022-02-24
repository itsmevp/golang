package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	http.Handle("/", LogMiddleWare(http.HandlerFunc(handler)))
	http.ListenAndServe("localhost:8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	println(r.Context(), "handler started")
	defer println(r.Context(), "handler ended")
}

func println(ctx context.Context, msg string) {
	if value, ok := ctx.Value("reqid").(int64); ok {
		fmt.Printf("[%d] - %s\n", value, msg)
	} else {
		log.Println("could not find the request ID")
		return
	}
}

func LogMiddleWare(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, "reqid", rand.Int63())
		f(w, r.WithContext(ctx))
	}
}
