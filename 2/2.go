package main

// sqrt(a + b) + sqrt(c + d)
import "fmt"
import "math"

func main() {

	var a float64 = 1
	var b float64 = 3
	var c float64 = 5
	var d float64 = 4
	var e float64
	e = math.Sqrt(a+b) + math.Sqrt(c+d)

	fmt.Print(e)

}
