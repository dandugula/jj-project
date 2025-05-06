package main

import "fmt"

// An awesome way to print hello world in Go
// Comment 1
func main() {
	print_hello()
	print_something()
	print_goodbye()
}

// Print hello
func print_hello() {
	fmt.Println("Hello World!")
}

func print_something() {
	fmt.Println("Something")
}

// Print goodbye
func print_goodbye() {
	fmt.Println("Goodbye World!")
}
