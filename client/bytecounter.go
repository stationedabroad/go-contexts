package main

import (
	"bytes"
	"bufio"
	"fmt"
	"io"
)

const input = `Now is the winter of our discontent,
			   Made glorious summer by this sun of York.
			   Now is the winter of our discontent,
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
	for word := 0; word < len(p); {
		adv, _, _ := bufio.ScanWords(p[word:], word != len(p))
		word += adv
		count++
	}
	*w += WordCounter(count)
	return count, nil
}

func (l *LineCounter) Write(p []byte) (int, error) {
	count := 0
	for line := 0; line < len(p); {
		adv, _, _ := bufio.ScanLines(p[line:], line != len(p))
		line += adv
		count++
	}
	*l += LineCounter(count)
	return count, nil
}

func CountingWriter(w io.Writer) (io.Writer, int64) {
	var bytesWritten int64
	var buf bytes.Buffer	
	bytesWritten, _ = buf.WriteTo(w)
	return &buf, bytesWritten
}

func main() {
	n := ByteCounter(10)
	l, _ := n.Write([]byte{110, 2})
	fmt.Println(n, l)

	var x WordCounter
	wordCount, _ := x.Write([]byte(input))
	fmt.Printf("word count for sentence :%d\nand w value :%v\n", wordCount, x)

	var y LineCounter
	// lineCount, _ := y.Write([]byte(input))
	// fmt.Printf("line count for sentence :%d\nand l value :%v\n", lineCount, y)

	fmt.Fprintf(&y, "%d", input)
	fmt.Println(y)

	// var z LineCounter
	wWriter, wwBytes := CountingWriter(&y)
	fmt.Printf("type: %T and type %T\n%v", wWriter, wwBytes)
}
