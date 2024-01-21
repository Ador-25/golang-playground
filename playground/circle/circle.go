package circle

import "math"

func Area(radius float64) float64 {
	return math.Pi * radius * radius
}

func Perimeter(radius float64) float64 {
	return math.Pi * 2 * radius
}
func GetAreaAndPerimeter(radius float64) (float64, float64) {
	return Area(radius), Perimeter(radius)
}
