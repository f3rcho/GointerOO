package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func (e *Employee) SetId(id int) {
	e.ID = id
}

func (e *Employee) GetId() int {
	return e.ID
}

func (e *Employee) SetName(name string) {
	e.Name = name
}

func (e *Employee) GetName() string {
	return e.Name
}

func main() {
	e := Employee{
		ID:        1,
		Name:      "John",
		Address:   "123 Main St",
		DoB:       time.Now(),
		Position:  "Developer",
		Salary:    50000,
		ManagerID: 1,
	}
	e.SetId(555)
	e.SetName("Jane555")
	fmt.Println(e.GetName())
	fmt.Println(e.GetId())
}

// func main() {
// 	m := make(map[string]int)
// 	m["a"] = 1
// 	fmt.Println(m["a"])

// 	s := []int{1, 2, 3}
// 	s = append(s, 8)
// 	for index, value := range s {
// 		fmt.Println(index, value)
// 	}
// 	c := make(chan int)
// 	go doSomething(c)
// 	<-c
// 	g := 25
// 	fmt.Println(g)
// 	h := &g
// 	fmt.Println(h)
// 	fmt.Println(*h)
// }

// func doSomething(c chan int) {
// 	time.Sleep(3 * time.Second)
// 	fmt.Println("Done")
// 	c <- 1
// }
