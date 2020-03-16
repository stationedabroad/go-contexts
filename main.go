package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		in := bufio.NewScanner(os.Stdin)
		in.Scan()
		fmt.Printf("cancelling value (just to check): %s\n", in.Text())
		cancel()
	}()

	mySleepAndTalk(ctx, 5 * time.Second, "Howdy Folks")
}

func mySleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	select {
	case <-time.After(d):
		fmt.Println(msg)
	case <-ctx.Done():
		log.Print(ctx.Err())
	}
}
