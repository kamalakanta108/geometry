package main

import "math"
import "fmt"

// s = r*r*p

func main() {

	var r float64 = 3
	var s float64
	s = r * r * math.Pi

	fmt.Print(s)

}
func Skr(r float64) float64 {
	return r * r * math.Pi
}
