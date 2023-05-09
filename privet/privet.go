package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// программа
func main() {
	fmt.Print("Введите длину: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	size, err := strconv.ParseFloat(input, 64)

	if err != nil {
		log.Fatal(err)
	}
	var status string

	if size <= 14 {
		status = "Повторяй Харе Кришна"
	} else if size == 15 {
		status = "Выбери духовного учителя"
	} else {
		status = "Старайся не общаться с противоположным полом"
	}

	fmt.Println(status)

}
