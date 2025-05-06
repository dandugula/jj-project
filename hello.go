package main

import "fmt"

// An awesome way to print hello world in Go

func main() {
	print_hello()
	print_goodbye()
}

// Print hello
func print_hello() {
	fmt.Println("Hello World!")
}

// Print goodbye
func print_goodbye() {
	fmt.Println("Goodbye World!")
}
