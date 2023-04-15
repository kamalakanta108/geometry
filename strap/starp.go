package main

import "fmt"

// s = (a + b)/2*h

func main() {
	var a float64 = 3
	var b float64 = 8
	var h float64 = 5
	var s float64
	s = Strp(a, b, h)
	fmt.Print(s)
}

// вычисление площади трапеции через сумму оснований [a+b] и высоту[h]
func Strp(a float64, b float64, h float64) float64 {
	return (a + b) * h / 2

}
