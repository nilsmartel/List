package main

import (
	"fmt"

	"./list"
)

func main() {
	list := list.New(0, 1, 2, 3, 4, 5)
	list.Push(6)
	list.Push("Seven")
	for {
		value := list.Pop()
		if value == nil {
			break
		}

		fmt.Println(value)
	}
}
