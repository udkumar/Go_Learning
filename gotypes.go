package main

import "fmt"

// func add(x float64, y float64) float64{
func add(x,y float64) float64{
	return x+y
}

func main(){
	//var num1, num2 float64 = 5.6, 9.5
	//var num2 float64 = 9.5
	num1,num2 := 5.6,9.5
	
	fmt.Print(add(num1,num2), "\n")
}
