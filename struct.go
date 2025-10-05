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
}
