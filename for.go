package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke", counter)
	// 	counter++
	// }

	// fmt.Println("Perulangan selesai")

	for counter := 1; counter <= 5; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	fmt.Println("Perulangan selesai")

	names := []string{"John", "Jane", "Doe"}
	for i := 0; i < len(names); i++ {
		fmt.Println("Index", i, "=", names[i])
	}

	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	}
	
	for _, name := range names {
		fmt.Println(name)
	}
}