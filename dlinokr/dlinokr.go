package main

import "math"
import "fmt"

// c = 2*P*r

func main() {

	var r float64 = 4
	var c float64
	c = 2 * r * math.Pi

	fmt.Print(c)

}
func Dlok(r float64) float64 {
	return 2 * r * math.Pi
}
