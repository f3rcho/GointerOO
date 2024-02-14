package main

import (
	"fmt"
)

func main() {
	tasks := []int{2, 3, 8, 9, 10, 30, 50}
	numWorkers := 3
	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))

	for i := 0; i < numWorkers; i++ {
		go Worker(i, jobs, results)
	}

	for _, task := range tasks {
		jobs <- task
	}
	close(jobs)

	for r := 0; r < len(tasks); r++ {
		<-results
	}
}

func Worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		fib := Fibonacci(j)
		fmt.Println("worker", id, "finished job", j, "and fib", fib)
		results <- fib
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
