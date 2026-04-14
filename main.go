package main

import (
	"fmt"
	"os"
)

func add(a, b int) int {
	return a + b
}

func main() {
	var a, b int
	fmt.Print("Enter two numbers: ")
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
}