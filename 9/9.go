package main

// a = b + c + d
import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var barg string = os.Args[1]
	var carg string = os.Args[2]
	var darg string = os.Args[3]
	var err error
	var b float64
	var c float64
	var d float64
	b, err = strconv.ParseFloat(barg, 64)
	if err != nil {
		panic(err)
	}
	c, err = strconv.ParseFloat(carg, 64)
	if err != nil {
		panic(err)
	}
	d, err = strconv.ParseFloat(darg, 64)
	if err != nil {
		panic(err)
	}
	var a float64
	a = A(b, c, d)
	fmt.Print(a)
}
func A(b float64, c float64, d float64) float64 {
	return b + c + d
}
