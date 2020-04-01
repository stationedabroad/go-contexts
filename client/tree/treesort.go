package tree

import "fmt"

type Tree struct {
    value       int
    left, right *Tree
}

func NewTree() *Tree {
    return &Tree{}
}

func (t *Tree) Value() int {
    return t.value
}

func (t *Tree) GetRight() *Tree {
    return t.right
}

func (t *Tree) GetLeft() *Tree {
    return t.left
}

func (t *Tree) PrintTree() string {
    
}

// Sort sorts values in place.
func (t *Tree) Sort(values []int) {
    var root *Tree
    for _, v := range values {
        root = add(t, v)
        fmt.Println(*root)
    }
    appendValues(values[:0], t)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *Tree) []int {
    if t != nil {
        values = appendValues(values, t.left)
        values = append(values, t.value)
        values = appendValues(values, t.right)
    }
    return values
}

func add(t *Tree, value int) *Tree {
    if t == nil {
        // Equivalent to return &Tree{value: value}.
        t = new(Tree)
        t.value = value
        fmt.Println("assign value", value)
        return t
    }
    if value < t.value {
        t.left = add(t.left, value)
    } else {
        fmt.Println("right hit", value)
        t.right = add(t.right, value)
    }
    return t
}
