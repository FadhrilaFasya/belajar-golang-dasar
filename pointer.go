package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// address1 := Address{
	// 	City:     "Jakarta",
	// 	Province: "DKI Jakarta",
	// 	Country:  "Indonesia",
	// }
	// address2 := address1 // Copying the value of address1 to address2

	// address2.City = "Bandung" // Changing the city in address2
	// fmt.Println("Address 1:", address1) // This will still show Jakarta
	// fmt.Println("Address 2:", address2) // This will show Bandung

	var address1 = Address{
		City:     "Jakarta",
		Province: "DKI Jakarta",
		Country:  "Indonesia",
	}
	var address2 = &address1 // Using a pointer to address1

	address2.City = "Bandung" // Changing the city through the pointer address2
	fmt.Println("Address 1:", address1) // This will still show Jakarta
	fmt.Println("Address 2:", address2) // This will show Bandung
}