package main

import "fmt"

func main() {
	name := "John Doe"

	switch name {
	case "John":
		fmt.Println("Hello John")	
	case "Jane":
		fmt.Println("Hello Jane")
	default:
		fmt.Println("Hello Stranger")
	}

	switch length := len(name); length > 5 {
		case true:
			fmt.Println("Name is too long")
		case false:
			fmt.Println("Name length is fine")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Name is too long")
	case length < 3:
		fmt.Println("Name is too short")
	default:
		fmt.Println("Name length is fine")
	}
}