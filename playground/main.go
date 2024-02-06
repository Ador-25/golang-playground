package main

import (
	"fmt"
	"math/rand"
	"playground/calculator"
	"playground/circle"
	thread "playground/goroutines"
	"strings"
	"time"
)

var gl_num int
var gl_bool bool

const e = 2.73

func getHeader(str string) {
	padding := "------------------------------"
	fmt.Println(padding, strings.ToUpper(str), padding)
}

func packageTest() {
	getHeader("package")
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
	getHeader("function")
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
	getHeader("variable")
	var python, java, c bool
	var i int
	fmt.Println(i, python, java, c)
	fmt.Println(gl_num, gl_bool)
	var j, k, l = 1, true, "str"
	fmt.Println(j, k, l)
	var f, g int = 3, 4
	fmt.Println(f, g)
}
func loopTest() {
	getHeader("loop")
	fmt.Println("for =>")
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	fmt.Println("while")
	itr := 1
	for itr <= 10 {
		fmt.Print(itr, " ")
		itr++
	}
}
func conditionals() {
	getHeader("conditionals")
	itr := 1
	for itr <= 10 {
		store := rand.Intn(100)
		if store%2 == 0 {
			fmt.Println(store, "is even")
			continue
		} else {
			fmt.Println(store, "is odd")
		}
		itr += 1
	}

	magicNumber := rand.Intn(500)
	fmt.Println("magic number =", magicNumber)
	switch magicNumber % 3 {
	case 1:
		fmt.Println("rem = 1")
	case 2:
		fmt.Println("rem = 2")
	default:
		fmt.Println("DIVISIBLE!!!!")
	}
}

func pointers() {
	i := 5
	j := &i
	*j++
	fmt.Println(i, *j)
}
func powers(x int) (int, int) {
	return x * x, x * x * x
}

func threadsAndChannels() {
	go thread.Say("external")
	thread.Say("main")
	c := make(chan int)
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("main array :=", arr)
	fmt.Println("channel :=", c)
	go thread.RangeSum(0, 2, arr, c)
	go thread.RangeSum(3, 4, arr, c)
	go thread.RangeSum(0, 4, arr, c)
	// these can be received in any order
	// check with sleep with random duration
	x := <-c
	y, z := <-c, <-c

	fmt.Println(x)
	fmt.Println(y, z)
	time.Sleep(100)
	close(c)
	//

	thread.BufferChannel(5)
	//time.Sleep(100)
	thread.BufferChannel(4)
	//time.Sleep(100)
	thread.BufferChannel(12)

	anotherChannel := make(chan int)
	//time.Sleep(1)
	go thread.AssignValue(anotherChannel)
	time.Sleep(time.Duration(rand.Int31n(100)))
	fmt.Println("value assigned = ", <-anotherChannel)
	fmt.Println("value assigned = ", <-anotherChannel)
	//fmt.Println("value assigned = ", <-anotherChannel)

}
func main() {
	x, y := 5, 8
	fmt.Println(x, y)
	threadsAndChannels()
}
