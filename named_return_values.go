package main

import "fmt"

func getFullName(firstName string, lastName string) string {
	fullName := firstName + " " + lastName
	return fullName
}

func main() {
	fmt.Println(getFullName("John", "Doe"))
}