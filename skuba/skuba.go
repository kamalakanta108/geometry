package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var aarg string = os.Args[1]
	var err error
	var a float64
	a, err = strconv.ParseFloat(aarg, 64)
	if err != nil {
		panic(err)
	}
	var s float64

	s = Skuba(a)
	fmt.Print(s)
}
func Skuba(a float64) float64 {
	return a * a * a
}
