package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

func main() {
	total := sumAll(10, 20, 30, 40, 50)
	fmt.Println("Total:", total)
}