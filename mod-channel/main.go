package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.Mutex
	i :=0
	for x :=0; x < 1000; x++{
		go func(){
			m.Lock()
			i++
			m.Unlock()
		}()
	}
	// wg.Wait()
	time.Sleep(2 * time.Second)
	fmt.Println(i)

}
