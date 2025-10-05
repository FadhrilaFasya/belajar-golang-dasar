package main

import "fmt"

type Customer struct {
	Name    string
	Address string
	Age     int
}

func main() {
	var cust Customer
	cust.Name = "Eko"
	cust.Address = "Subang"
	cust.Age = 30
	fmt.Println(cust)
	fmt.Println(cust.Name)
	fmt.Println(cust.Address)
	fmt.Println(cust.Age)

	john := Customer{
		Name:    "John",
		Address: "Jakarta",
		Age:     25,
	}
	fmt.Println(john)

	jane := Customer{"Jane", "Bandung", 23}
	fmt.Println(jane)
}
