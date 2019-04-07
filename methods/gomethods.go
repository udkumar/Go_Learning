package main

import(
	"fmt"
)

// two constants: 
const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
	gas_pedal uint16     //min: 0,      max: 65535    16bit
	brake_pedal uint16   //min: 0,      max: 65535
	steering_wheel int16 //min: -32768  max: 32768
	top_speed_kmh float64 //what's our top speed?
}

// we can use a method on a value, called a value receiver:
// gets a copy, receiver type

// you could just always use pointer receivers, but, value receivers are better for basic types and small structs. 
// This is because they're much cheaper to work with and can reduce the amount of overall garbage created.
//func (c car) kmh() float64 {

// Pointer reciever
func (c *car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh/usixteenbitmax)
}

// gets a copy, receiver type
func (c *car) mph() float64 {
	c.top_speed_kmh = 25 // 
	return float64(c.gas_pedal) * (c.top_speed_kmh/kmh_multiple/usixteenbitmax)
}

// we use a pointer receiver
func (c *car) new_top_speed(newspeed float64){
	c.top_speed_kmh = newspeed
}

// *car in func (c *car). Now, we're modifying the struct itself via pointer

 
func main() {
	//a_car := car{gas_pedal: 16535, brake_pedal: 0, steering_wheel: 12562, top_speed_kmh: 225.0}
	a_car := car{22314,0,12562,225.0}
	fmt.Println("gas_pedal:",a_car.gas_pedal, "brake_pedal:",a_car.brake_pedal,"steering_wheel:",a_car.steering_wheel)

	fmt.Println("Car is going",a_car.mph(),"MPH,",a_car.kmh(),"KMH, and top speed is",a_car.top_speed_kmh)
	a_car.new_top_speed(500)
	fmt.Println("Car is going",a_car.mph(),"MPH,",a_car.kmh(),"KMH, and top speed is",a_car.top_speed_kmh)
}

