// Go’s structs are typed collections of fields. They’re useful for grouping data together to form records.
package main

import "fmt"

//This person struct type has name and age fields.
type person struct {
	name string
	age int
}

func main() {
	//This syntax creates a new struct.
	fmt.Println(person{"uday", 30})

	//You can name the fields when initializing a struct.
	fmt.Println(person{name: "uday kumar", age: 30})

	//Omitted fields will be zero-valued.
	fmt.Println(person{name: "uday ji"})

	//An & prefix yields a pointer to the struct.
	fmt.Println(&person{name: "savita", age: 30})

	//Access struct fields with a dot.
	s := person{name: "ajita", age: 28}
	fmt.Println(s.name)

	//You can also use dots with struct pointers - the pointers are automatically dereferenced.
	sp := &s
    	fmt.Println(sp.age)

	//Structs are mutable.
	sp.age = 51
	fmt.Println(sp.age)
}

