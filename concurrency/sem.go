package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := make(chan int, 2)
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		c <- 1
		wg.Add(1)
		go doSomething(i, &wg, c)
	}
	wg.Wait()
}

func doSomething(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()

	fmt.Printf("Id %d start\n", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("Id %d end---\n", i)
	<-c
}
