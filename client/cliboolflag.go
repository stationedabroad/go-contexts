package main

import (
	"flag"
	"fmt"
)

var (
	someBool = flag.Bool("s", false, "select true or false (default:false)")
)

func main() {
	flag.Parse()
	fmt.Printf("parsed hook: %v\n", *someBool)
}