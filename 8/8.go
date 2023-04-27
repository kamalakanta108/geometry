package main

// d = (a+b)/(a-b)
import (
	"fmt"
	"os"
	"strconv"
)

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
	var d float64
	d = D(a, b)

	fmt.Print(d)

}
func D(a float64, b float64) float64 {
	return (a + b) / (a - b)
}
