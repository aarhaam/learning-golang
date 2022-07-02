package main

import "fmt"

func getFullname() (string, string, string) {
	return "go", "go.mod", "go.sum"
}

func main() {
	firstName, _, lastName := getFullname()

	fmt.Println(firstName)

	fmt.Println(lastName)
}
