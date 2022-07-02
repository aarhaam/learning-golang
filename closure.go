package main

import "fmt"

func main() {
	counter := 0
	name := "Rian"

	increment := func() {
		name := "Budi"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
