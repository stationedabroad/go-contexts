package main

import (
	"fmt"
	"net"
	"sync"
)

var wg sync.WaitGroup

func main() {
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			conn, err := net.Dial("tcp", fmt.Sprintf("scanme.nmap.org:%d", port))
			if err == nil {
				fmt.Printf("Open port: %d\n", port)
				conn.Close()
			}
		}(i)
	}
	wg.Wait()
}