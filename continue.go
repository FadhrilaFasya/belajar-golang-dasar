package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Skipping perulangan ke", i)
			continue
		}
		fmt.Println("Perulangan ke", i)
	}
}