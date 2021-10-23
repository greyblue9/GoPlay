package main

import (
	"fmt"
)

// Implemented Singly Linkedlist DS in golang

type Node struct {
	Value int
	Next  *Node
}

var root *Node
var size int

func addNodeAtEnd(t *Node, v int) int {

	if root == nil {
		t = &Node{v, nil}
		root = t
		size += 1
		return 0

	}

	if v == t.Value {
		fmt.Println("Node already exits")
		return -1
	}

	if t.Next == nil {
		t.Next = &Node{v, nil}
		size += 1
		return -2
	}

	return addNodeAtEnd(t.Next, v)

}

func addNodeAtFront(t *Node, val int) {
	if t == nil {
		fmt.Println("Empty List ")
	}

	t = &Node{val, nil}
	t.Next = root
	root = t

}

func Lsize(t *Node) int {
	if t == nil {
		fmt.Println("-> Empty List")
		return 0
	}

	i := 0
	for t != nil {
		i += 1
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
		fmt.Print(t.Value, " -> ")
		t = t.Next
	}
	print("\n")
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

func addNodeAtAny(t *Node, pos, v int) {
	if t == nil {
		fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
		fmt.Println("Emty List")
		fmt.Println("you can't traverse through an empty List")
	}

	abs := 1
	index := 2
	for t.Next != nil {

		temp := t.Next
		oldt := t
		if pos == abs {
			t = &Node{v, oldt}
			root = t
			return
		}

		if pos == index {
			t.Next = &Node{v, temp}
			return
		}

		t = t.Next
		index += 1

	}

}

func reverse(t *Node) *Node {
	if t == nil {
		fmt.Println("Noting To Reverse")
		return nil
	}

	cur := t
	var prev *Node = nil

	for cur != nil {
		prev, cur, cur.Next = cur, cur.Next, prev
	}

	return prev
}

func delNodeAtEnd(t *Node) {
	if t == nil {
		fmt.Println("Empty List")
	}

	temp := t
	for t.Next != nil {
		temp = t
		t = t.Next
	}

	temp = temp
	temp.Next = nil
}

func delNodeAtFront(t *Node) {
	if t == nil {
		fmt.Println("Empty List")
	}

	temp := t.Next
	t = nil
	root = temp

}

func delNodeAtAny(t *Node, pos int) {
	if t == nil {
		fmt.Println(" Empty List")
	}

	i := 0
	temp := t
	for t != nil {
		i += 1

		if pos == 1 {
			temp = t.Next
			root = temp
			tout := t
			tout = tout
			tout = nil
			return
		}

		if pos == i {
			temp.Next = t.Next
			t = nil
			return
		}

		//temp = temp
		temp = t
		t = t.Next
	}

}
