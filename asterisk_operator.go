package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{
		City:     "Jakarta",
		Province: "DKI Jakarta",
		Country:  "Indonesia",
	}
	var address2 *Address = &address1

	address2.City = "Bandung" 
	fmt.Println("Address 1:", address1)
	fmt.Println("Address 2:", address2)

	// address2 = &Address{
	// 	City:     "Malang",
	// 	Province: "Jawa Timur",
	// 	Country:  "Indonesia",
	// }
	*address2 = Address{
		City:     "Malang",
		Province: "Jawa Timur",
		Country:  "Indonesia",
	}
	fmt.Println("Address 1:", address1)
	fmt.Println("Address 2:", address2)
}