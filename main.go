package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func main() {
	var a, b int
	fmt.Print("Enter two numbers: ")
	fmt.Scan(&a, &b)
	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))
}