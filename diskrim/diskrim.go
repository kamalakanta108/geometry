package main

import "math"
import "fmt"

// diskrim = math.sqrt(b*b - 4*a*c)

func main() {
	var b float64 = 3
	var a float64 = 1
	var c float64 = -5
	var diskrim float64

	diskrim = math.Sqrt(b*b - 4*a*c)
	fmt.Print(diskrim)
}

func Diskrim(a float64, b float64, c float64) float64 {
	return math.Sqrt(b*b - 4*a*c)
}
