package main

import (
	"fmt"
	"bufio"
)

const input = `Now is the winter of our discontent,
			   Made glorious summer by this sun of York.
			   Now is the winter of our discontent,
			   Made glorious summer by this sun of York.`

type ByteCounter int
type WordCounter int
type LineCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func (w *WordCounter) Write(p []byte) (int, error) {
	count := 0
	for step := 0; step < len(p); {
		adv, _, _ := bufio.ScanWords(p[step:], step != len(p))
		step += adv
		count++
	}
	*w += WordCounter(count)
	return count, nil
}

func main() {
	n := ByteCounter(10)
	l, _ := n.Write([]byte{110, 2})
	fmt.Println(n, l)

	var x WordCounter
	wordCount, _ := x.Write([]byte(input))
	fmt.Printf("word count for sentence :%d\nand w value :%v\n", wordCount, x)
}
