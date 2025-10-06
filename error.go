package main

import "errors"

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagi tidak boleh nol")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	hasil, err := Pembagian(10, 0)
	if err != nil {
		println("Error:", err.Error())
	} else {
		println("Hasil:", hasil)
	}
}