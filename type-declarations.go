package main

import "fmt"

func main() {
	type NoKTP string
	type marriedStatus bool

	var noKtpEko NoKTP = "64871892738912"
	var marriedStatusMarried marriedStatus = false

	fmt.Println(noKtpEko)
	fmt.Println(marriedStatusMarried)
}
