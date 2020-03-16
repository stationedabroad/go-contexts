package main

import (
//	"bufio"
	"context"
	"fmt"
	"log"
//	"os"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
//	ctx, cancel := context.WithCancel(ctx)

	time.AfterFunc(time.Second, cancel)
//	go func() {
//		time.Sleep(time.Second)
//		cancel()
//	}()

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
