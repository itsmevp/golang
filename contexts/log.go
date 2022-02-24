package log

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
)

type logkey string

const tid logkey = "traceid"

func Println(ctx context.Context, msg string) {
	if reqid, ok := ctx.Value(tid).(int64); ok {
		fmt.Printf("[%d]-%s\n", reqid, msg)
	}
}

func LogMiddlerWare(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, tid, rand.Int63())
		f(w, r.WithContext(ctx))
	}
}
