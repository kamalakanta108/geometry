package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// hyp(l1, l2) -> sqrt(l1*l1 + l2*l2)
// float64

func main() {
	var aarg string = os.Args[1]
	var barg string = os.Args[2]
	var err error
	var a float64
	var b float64
	a, err = strconv.ParseFloat(aarg, 64)
	if err != nil {
		panic(err)
	}
	b, err = strconv.ParseFloat(barg, 64)
	if err != nil {
		panic(err)
	}
	var c float64
	c = Katet(a, b)
	fmt.Print(c)

}
func Katet(a float64, b float64) float64 {
	return math.Sqrt(a*a + b*b)
}
