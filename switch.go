package main

import "fmt"

func main() {

	name := "Eko"

	switch name {
	case "Eko":
		fmt.Println("Hello, Eko!")
	case "Joko":
		fmt.Println("Hello, Joko!")
	default:
		fmt.Println("Hello Dika")
	}

	//short statement switch
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
		break
	case false:
		fmt.Println("Nama sudah benar")
	}

}
