package main

import "fmt"

func main() {
	// Unlike arrays, slices are typed only by the elements they contain.
	// To create an empty slice with non-zero length, use the builtin make.
	// Here we make a slice of strings of length 3 (initially zero-valued).
	s := make([]string, 3)
	fmt.Println("emp:", s)
} 
