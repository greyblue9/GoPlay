package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

var root *Node

func main() {

	addNode(root, 001)
	addNode(root, 281)
	addNode(root, 456)
	fmt.Println("List Size ->", size(root))

	if lookupNode(root, 456) {
		fmt.Println("Node Exist")
	} else {
		fmt.Println("Node don't Exits!")
	}

	traverse(root)

}

func addNode(t *Node, v int) int {

	if root == nil {
		t = &Node{v, nil}
		root = t
		return 0

	}

	if v == t.Value {
		fmt.Println("Node already exits")
		return -1
	}

	if t.Next == nil {
		t.Next = &Node{v, nil}
		return -2
	}

	return addNode(t.Next, v)

}

func size(t *Node) int {
	if t == nil {
		fmt.Println("-> Empty List")
		return 0
	}

	i := 0
	for t != nil {
		i++
		t = t.Next
	}

	return i
}

func traverse(t *Node) {
	if t == nil {
		fmt.Println("-> Empty List")
		return
	}

	for t != nil {
		fmt.Println(t.Value)
		t = t.Next
	}

	fmt.Println()
}

func lookupNode(t *Node, v int) bool {

	if root == nil {
		t := &Node{v, nil}
		root = t
		return false
	}

	if v == t.Value {
		return true
	}

	if t.Next == nil {
		return false
	}

	return lookupNode(t.Next, v)
}

func delNode(t *Node, v int) {
	if t == nil {
		fmt.Println("-> Empty List")
	}

}
