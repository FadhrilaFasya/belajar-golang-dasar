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

	book := make(map[string]string)
	book["title"] = "Go Programming"
	book["author"] = "John Doe"
	book["publisher"] = "Tech Books Publishing"

	fmt.Println(book)

	delete(book, "publisher")

	fmt.Println(book)
}