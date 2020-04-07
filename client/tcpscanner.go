package main

import (
	"fmt"
	"net"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1024)
	for i := 1; i <= 1024; i++ {
		go dialHost(fmt.Sprintf("scanme.nmap.org:%d", i), &wg)
	}
	wg.Wait()
}

func dialHost(add string, wg *sync.WaitGroup) {
	defer wg.Done()
	conn, err := net.Dial("tcp", add)
	if err == nil {
		fmt.Printf("Open host/port: %s\n", add)
		conn.Close()
	}
}