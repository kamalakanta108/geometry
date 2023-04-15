package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var aarg string = os.Args[1]
	var barg string = os.Args[2]
	var carg string = os.Args[3]
	var darg string = os.Args[4]
	var err error
	var a float64 = 4
	var b float64 = 5
	var c float64 = 6
	var d float64 = 7
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
	d, err = strconv.ParseFloat(darg, 64)
	if err != nil {
		panic(err)
	}

	var p float64
	p = Per(a, b, c, d)
	fmt.Print(p)
}
func Per(a float64, b float64, c float64, d float64) float64 {
	return a + b + c + d
}
