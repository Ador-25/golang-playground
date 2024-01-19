package main

import (
	"fmt"
	"math"
	"math/rand"
	"playground/calculator"
	"time"
)

func main() {
	fmt.Println("Hello, Golang!{x$}")
	a := rand.Intn(5)
	b := rand.Intn(5)
	var pi = math.Pi
	fmt.Println("The time is", time.Now())
	fmt.Println(a, " ", b)
	sum := calculator.Add(a, b)
	var diff = calculator.Sub(a, b)
	fmt.Println(diff)
	fmt.Println(sum)
	fmt.Println(pi)
}
