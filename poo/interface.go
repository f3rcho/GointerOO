package main

import (
	"fmt"
)

type Person2 struct {
	Id   int
	Name string
	Age  int
}

type Employee2 struct {
	Person2
	Salary   int
	Position string
}

type FullTimeEmployee struct {
	Person2
	Employee2
}

type TempEmployee struct {
	Person2
	Employee2
	TaxRate int
}

type PrintInfo interface {
	getMessage() string
}

func (f FullTimeEmployee) getMessage() string {
	return fmt.Sprintf("%s is full-time employee", f.Name)
}
func (f TempEmployee) getMessage() string {
	return fmt.Sprintf("%s is Temporary employee", f.Name)
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.Name = "John"
	ftEmployee.Age = 30
	ftEmployee.Id = 234
	tEmployee := TempEmployee{}
	tEmployee.Name = "Tem"
	getMessage(ftEmployee)
	getMessage(tEmployee)
}
