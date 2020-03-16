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
	defer cancel()

	pauseAndPrint(ctx, 5 * time.Second, "Howdy Folks")
}

func pauseAndPrint(ctx context.Context, d time.Duration, msg string) {
	select {
	case <-time.After(d):
		fmt.Println(msg)
	case <-ctx.Done():
		log.Print(ctx.Err())
	}
}
