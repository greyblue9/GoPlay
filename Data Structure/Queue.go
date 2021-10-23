package main

import (
	"fmt"
)

// Implementing Queue DS in golang

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
		fmt.Println(size, "item added to the Queue")
		fmt.Println("Total items added in Queue:", size)
		return
	}

	t = &Node{data, nil}
	t.Next = head
	head = t
	size += 1
	fmt.Println("==================================")
	fmt.Println(size, "item added to the Queues")
	fmt.Println("Total items added in Queue:", size)

}

func Pop(t *Node) {
	if size == 0 {
		fmt.Println("-------------------------------------")
		fmt.Println("Empty Queue")
		fmt.Println("you can't pop from an empty Queue")
		return
	}

	temp := head
	for t.Next != nil {
		temp = t
		t = t.Next
	}

	temp.Next = nil
	size -= 1
	fmt.Println("#########################################")
	fmt.Println("poped", count, "items from the Queue")
	fmt.Println("Remaining", size, "items in the Queue")
}

func traverse(t *Node) {
	if t == nil {
		fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
		fmt.Println("Emty Queue")
		fmt.Println("you can't traverse through an empty Queue")
		return

	}

	fmt.Println("_______________________________")
	temp := t
	fmt.Println("all items in the Queue")
	for t.Next != nil {
		temp = t.Next
		fmt.Println(t.Data)
		t = t.Next
	}

	fmt.Println(temp.Data)
}
