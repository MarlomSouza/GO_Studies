package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(3)
	go callAPI(&wg)
	go callDB(&wg)
	go internalProcess(&wg)

	wg.Wait()

}

func callDB(wg *sync.WaitGroup){
	time.Sleep(1 * time.Second)
	fmt.Println("Call DB END")
	wg.Done()

}

func callAPI(wg *sync.WaitGroup){
	time.Sleep(4 * time.Second)
	fmt.Println("Call API END")
	wg.Done()
}

func internalProcess(wg *sync.WaitGroup)  {
	time.Sleep(2 * time.Second)
	fmt.Println("Call internallProcess END")
	wg.Done()
}