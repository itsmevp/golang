package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("127.0.0.1:9999", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	reqID := rand.Int63()
	ctx = context.WithValue(ctx, "reqid", reqID)

	printLog(ctx, "handler started")
	defer printLog(ctx, "handler ended")

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprint(w, "Hello World!")
	case <-ctx.Done():
		printLog(ctx, ctx.Err().Error())
		http.Error(w, ctx.Err().Error(), http.StatusInternalServerError)
	}
}

func printLog(ctx context.Context, msg string) {
	if reqID, ok := ctx.Value("reqid").(int64); ok {
		log.Printf("[%v] - %s", reqID, msg)
	}
}
