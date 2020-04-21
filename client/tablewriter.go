package main

import (
	"fmt"
	"os"

	tw "github.com/olekukonko/tablewriter"
)

func main() {
	data := [][]string{
		[]string{"A", "The Good", "500"},
		[]string{"B", "The Bad", "450"},
		[]string{"C", "The Ugly", "469"},
		[]string{"D", "The Gopher", "525"},
	}

	if len(os.Args[3]) > 0 {
		fmt.Println(os.Args[3])
	}
	
	table := tw.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Description", "Rating"})

	for i, v := range data {
		fmt.Printf("appended entry: (%d)\n", i)
		table.Append(v)
	}
	// table.AppendBulk(data[2:])
	table.Render()
}