package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	for i, _ := range make([]int, 1000) {
		go showMessage(strconv.Itoa(i))
	}

	time.Sleep(time.Duration(time.Hour.Seconds() * float64(5)))
}

func showMessage(message string) {
	fmt.Println(message)
}