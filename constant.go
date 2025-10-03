package main

import "fmt"

func main() {
	const (
		firstName string = "John"
		lastName  string = "Doe"
	)

	fmt.Println("First Name =", firstName)
	fmt.Println("Last Name =", lastName)
	
	// error
	// firstName = "Edo"
	// lastName = "Wibisono"
}