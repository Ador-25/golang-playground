package main

import (
	"fmt"
	"math/rand"
	"playground/calculator"
	"playground/circle"
	"time"
)

func main() {
	// io
	var read int
	fmt.Println(read)

	fmt.Println("Hello, Golang!{x$}")
	// package
	a := rand.Intn(5)
	b := rand.Intn(5)
	//var pi = math.Pi
	fmt.Println("The time is", time.Now())
	fmt.Println(a, " ", b)
	// custom package
	sum := calculator.Add(a, b)
	var diff = calculator.Sub(a, b)
	fmt.Println(diff)
	fmt.Println(sum)

	// returning multiple values from function
	radius := 5.265
	var area = circle.Area(radius)
	perimeter := circle.Perimeter(radius)
	fmt.Print("Radius = ", radius)
	fmt.Println("Area = ", area)
	fmt.Println("Perimeter = ", perimeter)
	ar, pr := circle.GetAreaAndPerimeter(radius)
	fmt.Println("Area,Perimeter = ", ar, " ,", pr)
	number := 10
	s, q := calculator.SquareAndQube(number)
	fmt.Println("square = ", s)
	fmt.Println("qube = ", q)
}
