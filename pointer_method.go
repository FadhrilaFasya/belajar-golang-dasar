package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
} // pointer receiver

func main() {
	john := Man{Name: "John"}
	john.Married()

	fmt.Println(john.Name) 
}