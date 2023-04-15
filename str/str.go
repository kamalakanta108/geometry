package main

import "fmt"

// a*h/2

func main() {
	var a float64 = 10
	var h float64 = 3
	var s float64
	s = Str(a, h)
	fmt.Print(s)
}

// вычисление площади треугольника[s] по основанию[a] и высоте[h]
func Str(a float64, h float64) float64 {
	return a * h / 2
}
