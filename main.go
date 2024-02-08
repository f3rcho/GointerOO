package main

import (
	"fmt"
	"time"
)

func main() {
	m := make(map[string]int)
	m["a"] = 1
	fmt.Println(m["a"])

	s := []int{1, 2, 3}
	s = append(s, 8)
	for index, value := range s {
		fmt.Println(index, value)
	}
	c := make(chan int)
	go doSomething(c)
	<-c
	g := 25
	fmt.Println(g)
	h := &g
	fmt.Println(h)
	fmt.Println(*h)
}

func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
	c <- 1
}
