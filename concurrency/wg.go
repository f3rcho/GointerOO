package main

import (
	"fmt"
	"sync"
	"time"
)

func main3() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doSomething2(i, &wg)
	}

	wg.Wait()
}

func doSomething2(u int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("doSomething %d\n", u)
	time.Sleep(time.Second * 1)
	fmt.Printf("End...%d\n", u)
}
