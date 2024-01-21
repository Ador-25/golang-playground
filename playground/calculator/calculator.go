package calculator

func Add(a int, b int) int {
	return a + b
}
func Sub(a int, b int) int {
	return a - b
}
func Mul(a, b int) int {
	return a * b
}
func Div(a, b int) float64 {
	return float64(a / b)
}

// naked return
func SquareAndQube(a int) (x int, y int) {
	x = a * a
	y = a * a * a
	return
}
