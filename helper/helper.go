package helper

var version = "1.0.0"
var Application = "Belajar Golang Dasar"

func sayGoodbye(name string) string {
	return "Goodbye " + name
} // private (tidak bisa diakses diluar package helper)

func SayHello(name string) string {
	return "Hello " + name
} // public (bisa diakses diluar package helper)