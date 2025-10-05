package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Terjadi kesalahan")
	}
	fmt.Println("Aplikasi berjalan dengan baik")
}

func main() {
	runApp(true)
}