package main

import "fmt"

func random() interface{} {
	return new
}

func main() {
	var result any = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	// var resultInt interface{} = result()
	// fmt.Println(reseultInt)

	switch value  := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Integer", value)
	default:
		fmt.Println("Unknown Type")
	}
}