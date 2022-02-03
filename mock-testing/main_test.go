package main

import "testing"

func TestGetFullTimeEmployee(t *testing.T) {
	table := []struct {
		id                 int
		dni                string
		mockFun            func()
		expectedFtEmployee FullTimeEmployee
	}{
		{
			id:  1,
			dni: "1",
			mockFun: func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{
						id:       1,
						Position: "CEO",
					}, nil
				}
				GetPersonByDNI = func(dni string) (Person, error) {
					return Person{
						DNI:  "1",
						name: "John Doe",
						edad: 24,
					}, nil
				}
			},
			expectedFtEmployee: FullTimeEmployee{
				Person: Person{
					DNI:  "1",
					name: "John Doe",
					edad: 24,
				},
				Employee: Employee{
					1,
					"CEO",
				},
			},
		},
	}
	//Se guardan las funciones orginales por si son necesarias en un futuro
	originalGetEmployeeById := GetEmployeeById
	originalGetPersonByDNI := GetPersonByDNI

	for _, test := range table {
		test.mockFun() //Inicializa los mock
		ft, err := GetFullTimeEmployee(test.id, test.dni)

		if err != nil {
			t.Errorf("Erro when getting Employee")
		}

		if ft.edad != test.expectedFtEmployee.edad {
			t.Errorf("Error with age, got %d expected %d", ft.edad, test.expectedFtEmployee.edad)
		}
	}

	//Se regresan las funciones a sus valores originales
	GetEmployeeById = originalGetEmployeeById
	GetPersonByDNI = originalGetPersonByDNI
}
