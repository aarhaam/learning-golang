package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Dika")
	data.PushBack("Joe")
	data.PushBack("Michael")
	data.PushFront("John")

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)

	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	fmt.Println("--------------------------------")

	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}
