package main

import (
	"fmt"

	"workspace/tree"
)

func main() {
	myTree := tree.NewTree()
	myTree.Sort([]int{-1,20, 30})
	fmt.Printf("New tree: %d, right: %d left: %d\n", myTree.Value(), myTree.GetRight().Value(), myTree.GetLeft().Value())
}