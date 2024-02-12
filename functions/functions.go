package main

import (
	"fmt"
	"time"
)

func main() {
	// anon functions
	x := 5
	y := func() int {
		return x * 2
	}()

	fmt.Println(y)
	// anon functions for go routines
	c := make(chan int)
	go func() {
		fmt.Println("Starting func...")
		time.Sleep(5 * time.Second)
		fmt.Println("Hello from a go routine")
		fmt.Println("End process!")
		c <- 1
	}()
	<-c
}
