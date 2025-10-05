package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi selesai")
	panicErr := recover()
	if panicErr != nil {
		fmt.Println("Terjadi error:", panicErr)
	}
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
	fmt.Println("Aplikasi tetap berjalan")
}