package main

import "fmt"

type Person struct {
	Id   int
	Name string
	Age  int
}

type Employee struct {
	Person
	Salary   int
	Position string
}

func (p Person) SayHello() {
	fmt.Printf("Hi, my name is %s\n", p.Name)
}

func (p Employee) SayHello() {
	fmt.Printf("Hi, I am a %s and my name is %s \n", p.Position, p.Name)
}

func main() {
	me := Person{Id: 666, Name: "Fernando", Age: 34}
	developer := Employee{
		Person:   Person{Id: 123, Name: "Adan", Age: 4},
		Salary:   120000,
		Position: "Dev",
	}

	me.SayHello()
	developer.SayHello()
}
