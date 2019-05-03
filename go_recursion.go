package main

import "fmt"

//This fact function calls itself until it reaches the base case of fact(0).
func fact(i int) int {
	if i == 0 {
		return 1
	}
	return i * fact(i-1)
}

func main() {
	fmt.Println(fact(3))
}
