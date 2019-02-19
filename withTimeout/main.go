package main

import (
	"context"
	"log"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func doSlow(ctx context.Context) error {
	defer wg.Done() // mark as done

	for i := 0; i < 10; i++ {
		t := 1 * time.Second
		select {
		case <-time.After(t):
			log.Println("Executing time...", i+1)
		case <-ctx.Done():
			log.Println("Cancel on timeout...", i+1)
			return ctx.Err()
		}
	}

	return nil
}

func main() {
	ctx := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	wg.Add(1) // number of go routine
	go doSlow(ctxWithTimeout)
	wg.Wait()
}
