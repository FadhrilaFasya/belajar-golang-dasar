package main

import "fmt"

func main() {

	type NoKTP string

	var ktpTest NoKTP = "111111111"
	var example string = "222222222"
	var exmpKtp NoKTP = NoKTP(example)

	fmt.Println(ktpTest)
	fmt.Println(exmpKtp)
}