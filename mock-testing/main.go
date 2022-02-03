package main

import "time"

type Person struct {
	DNI  string
	name string
	edad int
}

type Employee struct {
	id       int
	Position string
}

type FullTimeEmployee struct {
	Person   //Campo anonimo
	Employee //Campo anonimo
}

//Para que una funcion pueda ser mockeada debe ser declarada como variable
var GetPersonByDNI = func(dni string) (Person, error) {
	time.Sleep(5 * time.Second)
	//Select * from persona where ...
	return Person{}, nil
}

//Para que una funcion pueda ser mockeada debe ser declarada como variable
var GetEmployeeById = func(id int) (Employee, error) {
	time.Sleep(5 * time.Second)
	return Employee{}, nil
}

func GetFullTimeEmployee(id int, dni string) (FullTimeEmployee, error) {
	var ftEmployee FullTimeEmployee

	e, err := GetEmployeeById(id)
	if err != nil {
		return ftEmployee, err
	}
	ftEmployee.Employee = e
	p, err := GetPersonByDNI(dni)
	if err != nil {
		return ftEmployee, err
	}
	ftEmployee.Person = p
	return ftEmployee, nil
}
