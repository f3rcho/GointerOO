package main

import "testing"

func TestGetFullTimeEmployee(t *testing.T) {
	table := []struct {
		id               int
		Dni              string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id:  1,
			Dni: "12",
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
			expectedEmployee: FullTimeEmployee{
				Person: Person{
					Name: "John",
					Age:  30,
					Id:   1,
				},
				Employee: Employee{
					Dni:      "12",
					Position: "CEO",
				},
			},
		},
	}
	originalGetEmployeeByDNI := GetEmployeeByDNI
	originalGetPersonById := GetPersonById

	for _, test := range table {
		test.mockFunc()
		ft, err := GetFullTimeEmployeeById(test.id, test.Dni)
		if err != nil {
			t.Errorf("Error when getting full time employee: %s", err)
		}
		if ft.Age != test.expectedEmployee.Age {
			t.Errorf("Error when getting full time employee: expected age %d, got %d", test.expectedEmployee.Age, ft.Age)
		}
		if ft.Name != test.expectedEmployee.Name {
			t.Errorf("Error when getting full time employee: expected name %s, got %s", test.expectedEmployee.Name, ft.Name)
		}
		if ft.Id != test.expectedEmployee.Id {
			t.Errorf("Error when getting full time employee: expected id %d, got %d", test.expectedEmployee.Id, ft.Id)
		}
		if ft.Dni != test.expectedEmployee.Dni {
			t.Errorf("Error when getting full time employee: expected Dni %s, got %s", test.expectedEmployee.Dni, ft.Dni)
		}
		GetEmployeeByDNI = originalGetEmployeeByDNI
		GetPersonById = originalGetPersonById
	}
}
