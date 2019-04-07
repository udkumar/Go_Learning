package main

import(
	"fmt"
)


// we have structs, and you can have methods by defining methods on the structs.
// You can think of it like they're "associated" with the struc.

// we create a car type using a struct
type car struct {
	gas_padel uint16 //min: 0, max: 65535
	break_padel uint16 //min: 0, max: 65535
	steering_wheel int16 // min: -32768, max: 32768
	top_speed_kmh float64
}

func main(){
// Use above struct, we can do as below:
// a_car := car{gas_pedal: 16535, brake_pedal: 0, steering_wheel: 12562, top_speed_kmh: 225.0}

// can also abbreviate it by passing the parameters in order
a_car := car{22314,0,12562,225.0}
fmt.Println("gas_pedel:",a_car.gas_padel, "break_padel:",a_car.break_padel, "steering_wheel:",a_car.steering_wheel, "top_speed_kmh:",a_car.top_speed_kmh)
}

// Notice how we have defined a_car to be acar struct, and then we can reference the struct 
// fields with a dot, like access attributes in a traditional class


