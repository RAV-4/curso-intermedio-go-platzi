package inventarioclases

import "fmt"

//Composicion sobre herencia, una clase puede tener atributos de otra clase, pero debe poder ser completamente ignorado
type Person struct {
	name string
	edad int
}

type Employee struct {
	id int
}

//Composicion
type FullTimeEmployee struct {
	Person   //Campo anonimo
	Employee //Campo anonimo
}

//Ya que GO no tiene implementacion de interfaces de manera explicita se debe crear un metodo por clase
func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Full time Employee"
}

type TemporalyEmployee struct {
	Person
	Employee
	taxRate int
}

//Ya que GO no tiene implementacion de interfaces de manera explicita se debe crear un metodo por clase
func (tEmployee TemporalyEmployee) getMessage() string {
	return "Temporaly time Employee"
}

//Por ultimo crear un metodo de implementación
func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

//Creacion de la interface
type PrintInfo interface {
	getMessage() string
}

func herenciaPolimorfimo() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Juan"
	ftEmployee.edad = 25
	ftEmployee.id = 1
	fmt.Println(ftEmployee)

	tEmployee := TemporalyEmployee{}
	getMessage(tEmployee)
	getMessage(ftEmployee)
}
