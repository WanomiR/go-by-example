package main

import "fmt"

const SIZE = 15

type Node struct {
	Value int
	Next  *Node
}

type HashTable struct {
	Table map[int]*Node
	Size  int
}

func main() {

	t := make(map[int]*Node, SIZE)
	table := &HashTable{Table: t, Size: SIZE}

	fmt.Println("Number of bins: ", table.Size)

	for i := 0; i < 100; i++ {
		table.insert(i)
	}
	table.traverse()

}

func (h *HashTable) hashFunc(value int) int {
	return value % h.Size
}

func (h *HashTable) insert(value int) int {
	key := h.hashFunc(value)
	element := Node{Value: value, Next: h.Table[key]}
	h.Table[key] = &element
	return key
}

func (h *HashTable) traverse() {
	for key := range h.Table {
		fmt.Printf("Key: %2d;  ", key)
		if h.Table[key] != nil {
			node := h.Table[key]
			for node != nil {
				fmt.Printf("%d -> ", node.Value)
				node = node.Next
			}
			fmt.Println()
		}
	}
}

func (h *HashTable) lookup(value int) bool {
	key := h.hashFunc(value)
	if h.Table[key] != nil {
		node := h.Table[key]
		for node != nil {
			if node.Value == value {
				return true
			}
			node = node.Next
		}
	}
	return false
}
