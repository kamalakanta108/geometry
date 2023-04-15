package main

import (
	"fmt"
	"math"
)

// gradus = gradus* 180/pi
func main() {

	var rad float64 = 1
	var gradus float64

	gradus = Grad(rad)
	fmt.Print(gradus)
}
func Grad(rad float64) float64 {

	return (rad * 180 / math.Pi)
}
