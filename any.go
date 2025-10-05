package main

import "fmt"

func ups() interface{} {
	// return "Ups"
	// return true
	return 123
}

func main() {
	var kosong any = ups()
	fmt.Println(kosong)
}