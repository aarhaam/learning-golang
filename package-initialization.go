package main

import (
	"fmt"
	"go-fundamental/database"
)

func main() {
	result := database.GetDatabases()
	fmt.Println(result)
}
