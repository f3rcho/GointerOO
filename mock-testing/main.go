package main

import "time"

type Person struct {
	Id   int
	Name string
	Age  int
}

type Employee struct {
	Person
	Dni      string
	Salary   int
	Position string
}

type FullTimeEmployee struct {
	Employee
	Person
}

var GetPersonById = func(id int) (Person, error) {
	time.Sleep(5 * time.Second)
	return Person{}, nil
}

var GetEmployeeByDNI = func(dni string) (Employee, error) {
	time.Sleep(5 * time.Second)
	return Employee{}, nil
}

func GetFullTimeEmployeeById(id int, dni string) (FullTimeEmployee, error) {
	var fTEmployee FullTimeEmployee

	p, err := GetPersonById(id)
	if err != nil {
		return fTEmployee, err
	}

	e, err := GetEmployeeByDNI(dni)
	if err != nil {
		return fTEmployee, err
	}
	fTEmployee.Person = p
	fTEmployee.Employee = e

	return fTEmployee, nil
}
