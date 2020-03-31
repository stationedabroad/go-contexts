package main

import (
	"fmt"
	// "os"
	// "io"
)

const (
	Connect = 14
)

// var w io.Writer

func main() {
	// w = os.Stdout
	// f := w.(*os.File)
	// fmt.Printf("type of f is: %T\n", f)
	// fmt.Printf("type of w is: %T\n", w)
	// newToken(Connect)
	// fmt.Printf("main type: %T\n", Connect)
	do(10)
}

func do(v interface{}) {
	i, ok := v.([]uint8)
	fmt.Println(i)
	fmt.Println(ok)
}