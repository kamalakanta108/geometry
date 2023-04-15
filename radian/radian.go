package main

import "math"
import "fmt"

// import "os" оперативная система
// import "strconv" тоже что то важное
import "os"
import "strconv"

func main() {
	// вычисление радиана через градус, где радиан равен градусы деленные на 180 и умноженные на пи
	var gradusarg string = os.Args[1]
	var err error
	var gradus float64
	gradus, err = strconv.ParseFloat(gradusarg, 64)
	if err != nil {
		panic(err)
	}
	var rad float64
	rad = Rad(gradus)
	fmt.Print(rad)

}

func Rad(grad float64) float64 {
	return ((grad / 180) * math.Pi)
}
