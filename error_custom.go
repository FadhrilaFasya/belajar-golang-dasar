package main

import "fmt"

type validationError struct {
	message string
} // custom error type

func (v *validationError) Error() string {
	return v.message
} // implement the error interface

type notFoundError struct {
	message string
}

func (n *notFoundError) Error() string {
	return n.message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{message: "ID tidak boleh kosong"}
	} // return pointer to validationError

	if id != "John" {
		return &notFoundError{message: "Data tidak ditemukan"}
	} // return pointer to notFoundError

	return nil // no error

}

func main() {
	err := SaveData("Jihn", nil) // simulate not found error
	if err != nil {
		// if validationErr, ok := err.(*validationError); ok {
		// 	fmt.Println("Validation Error:", validationErr.Error())
		// } else if notFoundErr, ok := err.(*notFoundError); ok {
		// 	fmt.Println("Not Found Error:", notFoundErr.Error())
		// } else {
		// 	fmt.Println("Other Error:", err.Error())
		// }

		switch finalErr := err.(type) { 
		case *validationError:
			fmt.Println("Validation Error:", finalErr.Error()) // type assertion
		case *notFoundError:
			fmt.Println("Not Found Error:", finalErr.Error()) // type assertion
		default:
			fmt.Println("Other Error:", finalErr.Error()) // type assertion
		}
	}

}