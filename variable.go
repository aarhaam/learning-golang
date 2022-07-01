package main

import "fmt"

func main() {
	var name string

	name = "Mahardika"
	fmt.Println(name)

	name = "John"
	fmt.Println(name)

	var friendName = "Budi"
	fmt.Println(friendName)

	var age = 30
	fmt.Println(age)

	country := "Indonesia"
	fmt.Println(country)

	// many declaration variables
	var (
		firstName = "John"
		lastName  = "Doe"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
