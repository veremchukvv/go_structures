package main

import (
	"fmt"
)

func main() {
	list := &List{}

	list.Push(1)
	list.Push(2)
	list.Push(3)
	list.Push(4)
	list.Push(5)

	fmt.Println(list.ToArray())

	list.Peek()

	list.Pop()
	list.Peek()

	fmt.Println(list.ToArray())

}
