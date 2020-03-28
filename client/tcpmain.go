package main

import (
	"bufio"
	"fmt"
	"flag"
	"net"
	"os"
	"time"
)

func runEvery(d time.Duration, f func()) {
	for x := range time.Tick(d) {
		fmt.Printf("call at time: %v\n ", x)
		f()
	}
}

func makeCall() {
	conn, err := net.Dial("tcp", "192.168.1.224:8001")
	if err != nil {
		fmt.Println("Error", err.Error())
		os.Exit(1)
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)
}

func main() {
	duration := flag.Int("duration", 1, "duration in seconds")
	flag.Parse()
	runTime := time.Duration(*duration) * time.Second
	runEvery(runTime, makeCall)
}

