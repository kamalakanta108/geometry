package main

// e = a/b + c/d

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
	var a float64
	var b float64
	var c float64
	var d float64
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
	var e float64
	e = E(a, b, c, d)
	fmt.Print(e)
}
func E(a float64, b float64, c float64, d float64) float64 {
	return a/b + c/d
}
