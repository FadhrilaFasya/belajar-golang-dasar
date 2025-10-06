package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var alamat1 *Address = new(Address) // membuat pointer baru yang mengarah ke alamat kosong
	var alamat2 *Address = alamat1 // alamat2 mengarah ke alamat yang sama dengan alamat1

	alamat2.Country = "Indonesia" // mengubah value Country melalui pointer alamat2

	fmt.Println(alamat1) // alamat1 juga berubah karena mengarah ke alamat yang sama dengan alamat2
	fmt.Println(alamat2) // alamat2 menampilkan value yang telah diubah
}