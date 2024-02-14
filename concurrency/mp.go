package main

import "time"

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	d1 := time.Second * 4
	d2 := time.Second * 2
	go doSome(d1, c1, 1)
	go doSome(d2, c2, 2)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			println("received", msg1, "from c1")
		case msg2 := <-c2:
			println("received", msg2, "from c2")
		}
	}
}

func doSome(i time.Duration, c chan<- int, param int) {
	time.Sleep(i)
	c <- param
}
