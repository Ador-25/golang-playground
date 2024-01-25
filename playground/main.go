package main

import (
	"fmt"
	"math/rand"
	"playground/calculator"
	"playground/circle"
	"time"
)

var gl_num int
var gl_bool bool

const e = 2.73

func packageTest() {
	fmt.Println("----------------PACKAGE-------------------")
	a := rand.Intn(5)
	b := rand.Intn(5)
	//var pi = math.Pi
	fmt.Println("The time is", time.Now())
	fmt.Println(a, " ", b)
	// custom package
	fmt.Println()
	sum := calculator.Add(a, b)
	var diff = calculator.Sub(a, b)
	fmt.Println(diff)
	fmt.Println(sum)
}
func functionTest() {
	fmt.Println("----------------FUNCTION-------------------")
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
	fmt.Println(calculator.SquareAndQube(5))
}

func variableTest() {
	// variable declares
	fmt.Println("----------------VARIABLE-------------------")
	var python, java, c bool
	var i int
	fmt.Println(i, python, java, c)
	fmt.Println(gl_num, gl_bool)
	var j, k, l = 1, true, "str"
	fmt.Println(j, k, l)
	var f, g int = 3, 4
	fmt.Println(f, g)
}

func main() {
	fmt.Println("Hello, Golang!{x$}")
	packageTest()
	functionTest()
	variableTest()
	// pointers
	i := 5
	j := &i
	*j++
	fmt.Println(i, *j)

}
