package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// программа по определению формы угла

func main() {
	fmt.Print("Ведите данные синуса")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	index, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if index < 1 {
		status = "острый или тупой"
	}
	if index == 1 {
		status = "прямой"
	}
	fmt.Println(status)
}
