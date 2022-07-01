package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "John",
		"address": "1600 Amphitheatre Parkway",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book = make(map[string]string)
	book["title"] = "Belajar Go-lang"
	book["author"] = "Richard Martin"
	book["ups"] = "Salah"

	fmt.Println(book)

	delete(book, "ups")
	fmt.Println(book)
}
