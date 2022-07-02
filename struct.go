package main

import "fmt"

type Customer struct {
	Name    string
	Address string
	Age     int
}

func main() {
	var eko Customer
	eko.Name = "Eko"
	eko.Address = "Indonesia"
	eko.Age = 30

	fmt.Println(eko)

	budi := Customer{"Budi", "Jakarta", 30}
	fmt.Println(budi)
}
