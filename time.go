package main

import (
	"fmt"
	"time"
)

func main() {

	utc := time.Date(2020, 10, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println(utc)

	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "2020-10-01")
	fmt.Println(parse)
}
