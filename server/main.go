package main

import (
	"fmt"
	"net/http"
	"time"

	"workspace/tmp/contexts/log"
)

// Context passing in http call
func main() {
	http.HandleFunc("/", log.Decorate(handler))
	panic(http.ListenAndServe(":8002", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println(ctx, "handler started!")
	defer log.Println(ctx, "handler ended ...")

	select {
	case <- time.After(5 * time.Second):
		fmt.Fprintln(w, "server started with context ...")
	case <- ctx.Done():
		err := ctx.Err()
		log.Println(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
