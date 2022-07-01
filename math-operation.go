package main

import "fmt"

func main() {

	var result = 10 + +10
	fmt.Println(result)

	var (
		a = 10
		b = 10
		c = a + b
	)

	fmt.Println(c)

	var (
		negative = -100
		positive = 100
	)

	fmt.Println(negative, positive)

}
