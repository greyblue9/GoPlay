package main

import (
	"fmt"
)

// Implemented DoublyLinkedList DS in Golang

type DNode struct {
	prev *DNode
	Data interface{}
	next *DNode
}

var head *DNode
var lnode *DNode

func addNode(t *DNode, val interface{}) {
        if t == nil {
                t = &DNode{nil, val, nil}
                lnode = t
                head = t
                return
        }

        prev := lnode
        for t != nil {
                t.next = &DNode{prev, val, nil}
                t = t.next
                lnode = t
                return
        }

}



func addDNodeAtEnd(t *DNode, val interface{}) {
	if t == nil {
		fmt.Println("Empty List")
		return
	}

	prev := lnode
	for t != nil {
		lnode.next = &DNode{prev, val, nil}
		t = t.next
		lnode = t
		return
	}
}

func addDNodeAtFront(t *DNode, val interface{}) {
	if t == nil {
		fmt.Println("Empty List")
		return
	}

	temp := t
	next := t
	temp = &DNode{nil, val, next}
	next.prev = temp
	head = temp
}

func addDNodeAtAny(t *DNode, pos, val interface{}) {
	if t == nil {
		fmt.Println("To Insert Node you must Already have a list")
		return
	}

	abs := 1
	index := 1

	for t != nil {
		next := t
		pre := t.prev

		if pos == abs {
			t = &DNode{nil, val, next}
			next.prev = t
			head = t
			return
		}

		if pos == index {
			t = &DNode{pre, val, next}
			pre.next = t
			next.prev = t
			return
		}

		t = t.next
		index += 1

	}

}

func delDNodeAtEnd(t *DNode) {
	if t == nil {
		fmt.Println("Empty List")
		return
	}

	temp := t
	for t.next != nil {
		temp = t
		t = t.next
	}

	temp.next = nil

}

func delDNodeAtFront(t *DNode) {
	if t == nil {
		fmt.Println("Empty List")
		return
	}

	temp := t.next
	temp.prev = nil
	head = temp

}

func delDNodeAtAny(t *DNode, pos interface{}) {
	if t == nil {
		fmt.Println("Empty List")
		return
	}

	i := 0
	temp := t
	for t != nil {
		temp = t
		i += 1

		if pos == 1 {
			temp = t.next
			temp.prev = nil
			head = temp
			return
		}

		v := t
		if t.next == nil {
			v = t.prev
			v.next = nil
			return
		}

		p := t.prev
		p2 := t.next
		if pos == i {
			t.next = nil
			t.prev = nil
			p.next = p2
			p2.prev = p
			return
		}

		t = t.next
	}

}

func traverse(t *DNode) {
	if t == nil {
		fmt.Println("Empty List")
		return
	}

	for t != nil {
		fmt.Print(" <– ", "|", t.Data, "|", " –> ")
		t = t.next
	}

	println("\n")

}

func reverse(t *DNode) {
	if t == nil {
		fmt.Println("Empty List")
		return
	}

	pre := t
	for t != nil {
		pre = t
		t = t.next
	}

	for pre != nil {
		fmt.Print(" <– ", "|", pre.Data, "|", " –> ")
		pre = pre.prev
	}

	println("\n")
}

func lookUpNode(t *DNode, val interface{}) {
	if t == nil {
		fmt.Println("Empty List")
	}

	i := 0
	for t != nil {
		if t.Data != val {
			i += 1
		} else {
			i -= 1
		}
		t = t.next
	}

	if i == Lsize(head) {
		fmt.Println("Node Not Exist")
	} else {
		fmt.Println("Node Exist")
	}
}

func Lsize(t *DNode) int {
	if t == nil {
		fmt.Println("Empty List")
		return -1
	}

	i := 0
	for t != nil {
		i += 1
		t = t.next
	}
	return i

}
