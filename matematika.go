package main

import "fmt"

func main() {
	var a = 10
	var b = 20
	var c = a + b
	fmt.Println("Hasil Penjumlahan:", c)

	var i = 10
	i += 10 // i = i + 10
	fmt.Println("Hasil i:", i)

	i += 5 // i = i + 5
	fmt.Println("Hasil i:", i)

	var j = 1
	j++ // j = j + 1
	fmt.Println("Hasil j:", j)
	j++ // j = j + 1
	fmt.Println("Hasil j:", j)
}