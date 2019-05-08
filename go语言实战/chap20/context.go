package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	context.WithTimeout()
	context.WithDeadline()
	go watch(ctx, "监控1")
	go watch(ctx, "监控2")
	go watch(ctx, "监控3")

	time.Sleep(10 * time.Second)

	cancel()
	time.Sleep(5 * time.Second)

}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "监控退出了...")
		default:
			fmt.Println(name, "goroutine 监控中")
			time.Sleep(2 * time.Second)
		}
	}
}
