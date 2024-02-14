package main

import "fmt"

// import "time"

// func ProduceNumbers(ch chan<- int) {
// 	for i := 0; i < 10; i++ {
// 		ch <- i
// 		ch <- i
// 	}
// 	close(ch)
// }

// func ConsumeNumbers(ch <-chan int) {
// 	for num := range ch {
// 		println(num)
// 	}
// }

// func main() {
// 	ch := make(chan int) // unbuffered channel
// 	go ProduceNumbers(ch)
// 	time.Sleep(5 * time.Second)
// 	ConsumeNumbers(ch)
// }

func main2() {
	ch := make(chan string, 3) // Create a buffered channel with capacity 3

	go func() {
		ch <- "message 1"
		ch <- "message 2"
		ch <- "message 3"
		close(ch)
	}()

	for msg := range ch {
		fmt.Println("Received:", msg)
	}
}
