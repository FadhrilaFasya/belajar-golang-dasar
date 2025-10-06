package main

import (
	"belajar-golang-dasar/database"
	"fmt"
)

func main() {
	fmt.Println("Database:", database.GetDatabase())
}