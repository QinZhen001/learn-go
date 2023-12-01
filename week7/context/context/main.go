package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	baseCtx := context.Background()
	ctx := context.WithValue(baseCtx, "key", "value")
	go func(c context.Context) {
		fmt.Println(c.Value("key"))
	}(ctx)
	timeoutCtx, cancel := context.WithTimeout(baseCtx, 3*time.Second)
	defer cancel()
	go func(ctx context.Context) {
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-ctx.Done():
				fmt.Println("timeout")
				return
			default:
				fmt.Println("work")
			}
		}
	}(timeoutCtx)

	<-timeoutCtx.Done()
	time.Sleep(1 * time.Second)
	fmt.Println("main process exit!")
}
