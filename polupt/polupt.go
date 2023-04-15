package main

import (
	"fmt"
	"os"
	"strconv"
)

// poluoerimetr p = (a+b+c)/2

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

	var p float64
	p = P(a, b, c)
	fmt.Print(p)
}

func P(a float64, b float64, c float64) float64 {
	return (a + b + c) / 2
}
