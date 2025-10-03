package main

import "fmt"

func main() {
	name := "John Doe"

	if name == "John" {
		fmt.Println("Hello John")
	} else if name == "Jane" {
		fmt.Println("Hello Jane")
	} else {
		fmt.Println("Hello Stranger")
	}
	
	if length := len(name); length > 5 {
		fmt.Println("Name is too long")
	} else {
		fmt.Println("Name length is fine")
	}

}