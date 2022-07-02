package main

import "fmt"

func getGoodBye(name string) string {
	return "goodbye " + name
}

func main() {

	sayGoodBye := getGoodBye
	result := sayGoodBye("Mahardika")

	fmt.Println(result)

}
