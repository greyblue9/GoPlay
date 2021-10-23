package main

import (
	"fmt"
)

// Implementing Stack DS in golang

type Node struct {
	Data int
	Next *Node
}

var head *Node
var size = 0
var count = 0

func Push(t *Node, data int) {
	if head == nil {
		head = &Node{data, nil}
		size += 1
		fmt.Println(size, "item added to the stack")
		fmt.Println("Total items added in Stack:", size)
		return
	}

	t = &Node{data, nil}
	t.Next = head
	head = t
	size += 1
	fmt.Println("==================================")
	fmt.Println(size, "item added to the stack")
	fmt.Println("Total items added in stack:", size)

}

func Pop(t *Node) {
	if size == 0 {
		fmt.Println("-------------------------------------")
		fmt.Println("Emty stack")
		fmt.Println("you can't pop from an empty stack")
		return
	}

	temp := t.Next
	t = nil
	head = temp
	size -= 1
	count += 1
	fmt.Println("#########################################")
	fmt.Println("poped", count, "items from the stack")
	fmt.Println("Remaining", size, "items in the stack")

}

func traverse(t *Node) {
	if t == nil {
		fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
		fmt.Println("Emty stack")
		fmt.Println("you can't traverse through an empty stack")
		return

	}

	fmt.Println("_______________________________")
	temp := t
	fmt.Println("all items in the stack")
	for t.Next != nil {
		temp = t.Next
		fmt.Println(t.Data)
		t = t.Next
	}

	fmt.Println(temp.Data)
}
