package main

import "fmt"

func main() {
	var month = [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	var slice1 = month[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	var slice2 = month[2:4]
	fmt.Println(cap(slice2))
	fmt.Println(slice2)

	var slice3 = append(slice2, "Dika")
	slice3[0] = "Bukan Maret"
	fmt.Println(slice3)
	fmt.Println(month)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Dika"
	newSlice[1] = "Muhammad"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	//copy slice

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)
}
