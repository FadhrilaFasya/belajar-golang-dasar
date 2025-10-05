package main

import "fmt"

func getGoodbye(name string) string {
	return "Goodbye " + name
}

func main() {
	example := getGoodbye

	fmt.Println(example("John"))
	fmt.Println(example("Doe"))
}