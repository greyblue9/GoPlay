package main

import (
	"fmt"
)

// Simple array DS in Golang

var arr [10]int
var pos = 0
var count = 0

func add(data int) {
	if pos >= 10 {
		fmt.Println("=============================================")
		fmt.Println("array is full -> You can only add 10 items")
		return
	}
	arr[pos] = data
	pos += 1
	fmt.Println("===============================")
	fmt.Println(pos, "item added to the array")
	fmt.Println("Total items added in array: ", pos)

}

func remove(index int) {
	for i := index + 1; i < len(arr); i++ {
		arr[i-1] = arr[i]
		arr[i] = 0
	}

	pos -= 1
	count += 1
	fmt.Println("########################################")
	fmt.Println(count, "item removed from the array")
	fmt.Println("Remaining", pos, "items in the array")
}

func traverse() {
	fmt.Println("************************")
	fmt.Println("All items in array")
	for _, v := range arr {
		fmt.Print(v, ",")
	}
	fmt.Print("\n")

}

func reverse() {
	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")
	fmt.Println("All items in array | reverse output")
	for i := 10 - 1; i >= 0; i-- {
		fmt.Print(arr[i], ",")
	}
	fmt.Print("\n")
}

func change(index, val int) {
	fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
	arr[index] = val
	fmt.Println("changed item in position", index, "to", val)
}
