package main

import "fmt"

func main() {

	var name = "Dika"

	if name == "Eko" {
		fmt.Println("Hello Eko!")
	} else if name == "Joko" {
		fmt.Println("Hello Joko!")
	} else {
		fmt.Println("Hi, Boleh Kenalan?")
	}

	var length = len(name)
	if length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

}
