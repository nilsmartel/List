package main

import (
	"fmt"

	"./list"
)

func main() {
	list := list.New(0, 1, 2, 3, 4, 4.5, 5)
	list.Push(6)
	list.Push("Seven")

	for !list.IsEmpty() {
		value := list.Pop()

		fmt.Println(value)
	}
}
