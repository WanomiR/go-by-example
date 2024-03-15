package main

import (
	"fmt"
	"math/rand"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func main() {
	tree := create(10)
	fmt.Println("Root node value: ", tree.Value)

	traverse := traverseWithCount()

	tree = insert(tree, -1)
	tree = insert(tree, -5)

	traverse(tree)
}

func traverse(t *Tree) {
	if t == nil {
		return
	}
	traverse(t.Left)
	fmt.Print(t.Value, " ")
	traverse(t.Right)
}

func traverseWithCount() func(t *Tree) {
	var cnt int
	var traverse func(t *Tree)

	traverse = func(t *Tree) {
		if t == nil {
			return
		}
		traverse(t.Left)
		cnt++
		fmt.Printf("%d: %d\n", cnt, t.Value)
		traverse(t.Right)
	}

	return traverse
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}

	if v == t.Value {
		return t
	}

	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

func create(n int) *Tree {
	var t *Tree
	rand.Seed(42)

	for i := 0; i < 2*n; i++ {
		val := rand.Intn(n * 2)
		t = insert(t, val)
	}

	return t
}
