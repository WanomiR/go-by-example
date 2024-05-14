package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func main() {
	tree := create(10)
	fmt.Println("Root node value: ", tree.Value)
	traverse(tree)
	fmt.Println("Max depth: ", findMaxDepth(tree))
}

func findMaxDepth(root *Tree) int {
	if root == nil {
		return 0
	}

	left := findMaxDepth(root.Left)
	right := findMaxDepth(root.Right)

	if left > right {
		return left + 1
	}

	return right + 1
}

func deepestLeavesSum(node *Tree, depth, maxDepth, leavesSum int) int {
	if node == nil {
		return 0
	}

	if depth == maxDepth {
		return leavesSum + node.Value
	}

	left := deepestLeavesSum(node.Left, depth+1, maxDepth, leavesSum)
	right := deepestLeavesSum(node.Right, depth+1, maxDepth, leavesSum)

	return left + right
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
	rand.Seed(time.Now().Unix())

	for i := 0; i < 2*n; i++ {
		val := rand.Intn(n * 2)
		t = insert(t, val)
	}

	return t
}
