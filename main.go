package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	ctx := context.Background()

	pauseAndPrint(ctx, 5 * time.Second, "Howdy Folks")
}

func pauseAndPrint(ctx context.Context, d time.Duration, msg string) {
	select {
	case t := <-time.After(d):
		   fmt.Printf("%s - time is: %v\n", msg, t)
	case <-ctx.Done():
		log.Print(ctx.Err())
	}
}
