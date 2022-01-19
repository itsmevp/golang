package main

import (
	"context"
	"fmt"
	"time"
)

/*
context deadline exceeded - cancelled after timeout duration
context canceled - cancelled explicitly
*/

func main() {
	ctx := context.Background()
	// ctx, cancel := context.WithCancel(ctx)
	// ctx, cancel := context.WithDeadline(ctx, <-time.After(1*time.Hour))
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	sleepAndTalk(ctx, 5*time.Second, "Hello World!")
}

func sleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	select {
	case time := <-time.After(d):
		fmt.Printf("%v:%v", time, msg)
	case <-ctx.Done():
		fmt.Println(ctx.Err().Error())
	}
}
