package main

import "fmt"

func main() {
	channel := make(chan int, 100) // BUFFER
	go setList(channel)
	
	for v := range channel {
		fmt.Println("reading", v)
	}
}
// setList(channel <-chan int)  <- only read
// setList(channel chan<- int)  <- insert
func setList(channel chan<- int) {

	for i := 0; i < 100; i++ {
		channel <- i
		fmt.Println("insert", i)
	}
	close(channel)
}

