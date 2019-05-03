package main

import "fmt"

func main() {
	// Unlike arrays, slices are typed only by the elements they contain.
	// To create an empty slice with non-zero length, use the builtin make.
	// Here we make a slice of strings of length 3 (initially zero-valued).
	s := make([]string, 3)
	fmt.Println("emp: ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("Set: ", s)
	fmt.Println("Get: ", s[2])

	//len returns the length of the slice as expected.
	fmt.Println("Len: ",len(s))

	//slices support several more that make them richer than arrays.
	//One is the builtin append, which returns a slice containing one or more new values.
	//Note that we need to accept a return value from append as we may get a new slice value.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append: ", s)

	// Here we create an empty slice c of the same length as s and copy into c from s.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy: ", c)

	//Slices support a “slice” operator with the syntax slice[low:high].
	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2: ", l)

	l = s[3:]
	fmt.Println("sl3: ", l)

	t := []string{"g","h","i"}
	fmt.Println("dcl: ", t)

} 
