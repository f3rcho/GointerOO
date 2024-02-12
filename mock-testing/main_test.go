package main

import "testing"

func TestGetFullTimeEmployee(t *testing.T) {
	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id:  1,
			dni: "12",
			mockFunc: func() {
				GetEmployeeByDNI = func(dni string) (Employee, error) {
					return Employee{
						Dni:      "12",
						Position: "CEO",
					}, nil
				}
				GetPersonById = func(id int) (Person, error) {
					return Person{
						Name: "John",
						Age:  30,
						Id:   1,
					}, nil
				}
			},
		},
	}
}
