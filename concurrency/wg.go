package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doSomething(i, &wg)
	}

	wg.Wait()
}

func doSomething(u int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("doSomething %d\n", u)
	time.Sleep(time.Second * 1)
	fmt.Printf("End...%d\n", u)
}
