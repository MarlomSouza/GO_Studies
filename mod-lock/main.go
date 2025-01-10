package main

import "fmt"

func main() {
	channel := make(chan int)
	go setList(channel)
	go setList2()
	for v := range channel {
		fmt.Println(v)
	}
}

func setList(channel chan int) {
	for i := 0; i < 100; i++ {
		channel <- i
	}
	close(channel)
}

func setList2() {
	for i := 100; i >=0 ; i-- {
		fmt.Println("Set list 2", i)
	}
	
}
