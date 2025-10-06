package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
) // import package (menggunakan package helper

func main() {
	result := helper.SayHello("John")
	fmt.Println(result)

	fmt.Println(helper.Application)
	// fmt.Println(helper.sayGoodbye("John")) // error, karena function sayGoodbye diawali dengan huruf kecil (private)
}