package main

import (
	"fmt"
	"time"
	"math/rand"
)

func randTime(){
	fmt.Println("Current time:-", time.Now())
	fmt.Println("Number print between 1-100:-", rand.Intn(100))
}

func main(){
	randTime()
}
