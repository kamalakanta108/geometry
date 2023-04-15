package main

import (
	"fmt"
	"os"
	"strconv"
)

// (a/b) + c
func main() {
	var aarg string = os.Args[1]
	var barg string = os.Args[2]
	var carg string = os.Args[3]
	var err error
	var a float64
	var b float64
	var c float64

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
	var d float64

	d = D(a, b, c)
	fmt.Print(d)
}
func D(a float64, b float64, c float64) float64 {
	return (a / b) + c
}
