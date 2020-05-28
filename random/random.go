package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	rand.Seed(time.Now().UnixNano())
	randomNumber := random(0, 10)
	fmt.Printf("The random number is: %d \n", randomNumber)
}
