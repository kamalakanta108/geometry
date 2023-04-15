package geometry

import "math"
import "fmt"

func main() {
	// вычисление квадратных уравнений 0 = a * x * x + b * x +c * x
	var a float64 = 1
	var b float64 = 10
	var c float64 = 39

	x1 = (-1*b + math.Sqrt(b*b-4*a*c)) / 2 * a
	x2 = (-1*b - math.Sqrt(b*b-4*a*c)) / 2 * a
	fmt.Print(x1, x2)
}

func QQ(x1 float64, x2 float64) float64 {

	return (-1*b + math.Sqrt(b*b-4*b*c)) / 2, (-1*b - math.Sqrt(b*b-4*b*c)) / 2
}
func Diskr(d float64) float64
{ 	return      (math.Sqrt(b*b - 4*a*c))}
