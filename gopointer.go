package main

import "fmt"

// Pointers allow you to reference from memory
func main(){
	x := 15
	// If we want to point to x, we can use the & symbol
	a := &x // point to x  (memory address)

	fmt.Println(a) // prints out the mem addr.

	// We can read through pointers using * before the variable
	fmt.Println(*a) // read a through the pointer, so this will print out a value (15 in this case)

	// The a variable was just a memory address, and the asterisk
	// gets us to read through to the memory address, when we combine these two,
	// we can actually modify the value at that memory address
	*a = 5  // sets the value pointed at to 5, which means x is modified (since x is stored at the mem addr)

	fmt.Println(x) // see the new value of x

	*a = *a**a // what is the value of x now?
	fmt.Println(x) // prints a value
	fmt.Println(*a) // prints a value

	// how might you get the memory address using x and a?
	fmt.Println(&x) // prints a memory address
	fmt.Println(a) // prints a memory address
}
