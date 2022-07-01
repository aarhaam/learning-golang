package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "A"
	names[1] = "B"
	names[2] = "C"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	fmt.Println(len(names))
}
