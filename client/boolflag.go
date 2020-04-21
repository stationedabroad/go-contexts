package main

import (
	"flag"
	"fmt"
)

var b bool

func main() {
	fs := flag.NewFlagSet("Example Bool", flag.ExitOnError)
	fs.BoolVar(&b, "p", false, "enter a true or false")
	fs.Parse([]string{"-p", "true"})
	fmt.Printf("the value of flag is %v\n", b)

	s := "hello, world"
	hello := s[:5]
	world := s[5:]
	ns := hello + world

	fmt.Printf("addr of s: %v - addr of hw: %v \n", &s, &ns)
}