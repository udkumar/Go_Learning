package main

import (
	"fmt"
)

// If function return two value, you specify their types inside parenthesis
func multiple(a,b string) (string, string){
	return a,b
}

func main(){
	w1,w2 := multiple("hey", "strange Go")
	fmt.Println(w1,w2)
	
	//If you wanted to convert the types of a variable
	var a int = 2
	var b float64 = float64(a)
	fmt.Println(b)

	var x float32 = 3
	y := x // y is float32 type
	fmt.Println(y)
}
