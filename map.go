package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}
	// person["name"] = "John"
	// person["age"] = "30"
	// person["city"] = "New York"

	person := map[string]string{
		"name": "John",
		"age":  "30",
		"city": "New York",
	}

	fmt.Println(person["name"])
	fmt.Println(person["age"])
	fmt.Println(person["city"])
	fmt.Println(person)
}