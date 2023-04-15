package main

import (
	"fmt"
	"os"
	"strconv"
)

// s = (a*b*c)/4*R  math
func main() {
	var aarg string = os.Args[1]
	var barg string = os.Args[2]
	var carg string = os.Args[3]
	var rarg string = os.Args[4]
	var err error
	var a float64 = 4
	var b float64 = 3
	var c float64 = 2
	var r float64 = 2
	a, err = strconv.ParseFloat(aarg, 64)
	if err != nil {
		panic(err)
	}
	b, err = strconv.ParseFloat(barg, 64)
	if err != nil {
		panic(err)
	}
	c, err = strconv.ParseFloat(carg, 64)
	if err != nil {
		panic(err)
	}
	r, err = strconv.ParseFloat(rarg, 64)
	if err != nil {
		panic(err)
	}

	var s float64
	s = Str(a, b, c, r)
	fmt.Print(s)
}
func Str(a float64, b float64, c float64, r float64) float64 {

	return (a * b * c) / (r * 4)
}
