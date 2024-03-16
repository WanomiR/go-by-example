package main

import "fmt"

type Node struct {
	Value    int
	Previous *Node
	Next     *Node
}

func main() {
	root := &Node{}
	fmt.Println(root)
	traverse(root)

	root = initRoot(root, 1)

	addNode(root, 1)
	traverse(root)

	addNode(root, 10)
	addNode(root, 5)
	addNode(root, 8)
	addNode(root, -2)

	fmt.Println("List size: ", size(root))
	traverse(root)
	reverse(root)
}

func initRoot(root *Node, value int) *Node {
	node := &Node{value, nil, nil}
	root = node
	return root
}

func addNode(node *Node, value int) int {
	if value == node.Value {
		fmt.Println("Node already exists: ", value)
		return -1
	}

	if node.Next == nil {
		tmp := node
		node.Next = &Node{value, tmp, nil}
		return -2
	}

	return addNode(node.Next, value)
}

func traverse(node *Node) {
	if node == nil {
		fmt.Println("-> Empty list!")
	}

	for node != nil {
		fmt.Printf("%d -> ", node.Value)
		node = node.Next
	}
	fmt.Println()
}

func reverse(node *Node) {
	if node == nil {
		fmt.Println("-> Empty list!")
	}

	reverseNode := node
	for node != nil {
		reverseNode, node = node, node.Next
	}

	for reverseNode.Previous != nil {
		fmt.Printf("%d -> ", reverseNode.Value)
		reverseNode = reverseNode.Previous
	}
	fmt.Printf("%d -> ", reverseNode.Value)
	fmt.Println()
}

func size(node *Node) int {
	if node == nil {
		fmt.Println("-> Empty list!")
	}

	cnt := 0
	for node != nil {
		cnt++
		node = node.Next
	}

	return cnt
}

func lookupNode(node *Node, value int) bool {
	if value == node.Value {
		return true
	}

	if node.Next == nil {
		return false
	}

	return lookupNode(node.Next, value)
}
