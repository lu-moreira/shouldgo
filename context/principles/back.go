package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

/*
context.Context, provides those two principal things:
    - cancellation propagation
    - send values


context.Background() --> used as rooted context
*/

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		time.Sleep(1 * time.Second)
		cancel()
	}()
	sleepAndTalk(ctx, 2*time.Second, "hello")
}

func sleepAndTalk(ctx context.Context, wait time.Duration, msg string) {
	select {
	case <-time.After(wait):
		fmt.Println(msg)
	case <-ctx.Done():
		log.Println(ctx.Err())
	}
}
