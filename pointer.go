package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1

	address2.City = "Bandung"

	address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesa"}

	fmt.Println(address1)
	fmt.Println(address2)

	var address4 *Address = new(Address)
	address4.City = "Jakarta"
	fmt.Println(address4)

	var alamat = Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "",
	}

	ChangeCountryToIndonesia(&alamat)
	fmt.Println(alamat)
}
