package inventarioclases

import "fmt"

//en este curso de toma como premisa que GO si es un lenguaje orientado a objetos
//Structs vs Clases

type claseEmpleado struct {
	id   int
	name string
}

//Constructor
func newEmployee(id int, name string) *claseEmpleado {
	return &claseEmpleado{
		id:   id,
		name: name,
	}
}

func (e *claseEmpleado) SetId(id int) {
	e.id = id
}

func (e *claseEmpleado) getId() int {
	return e.id
}

func (e *claseEmpleado) setName(name string) {
	e.name = name
}

func (e *claseEmpleado) getName() string {
	return e.name
}

func claseConstructores() {
	e := claseEmpleado{}
	fmt.Println(e)
	e.id = 1
	e.name = "Name"
	fmt.Println(e)
	e.SetId(5)
	e.setName("Nombre")
	fmt.Println(e.getId())
	fmt.Println(e.getName())

	//Formas de instanciar un struct
	//1
	e1 := claseEmpleado{}
	fmt.Printf("Forma 1 %v\n", e1)

	//2
	e2 := claseEmpleado{
		id:   10,
		name: "Empleado forma 2",
	}
	fmt.Printf("Forma 2 %v\n", e2)

	//3
	e3 := new(claseEmpleado) //Esta forma te retorna es la referencia (apuntador) a la instancia struct
	fmt.Printf("Forma 3: apuntador: %v\n", e3)
	fmt.Printf("Forma 3: valor: %v\n", *e3)

	//4
	e4 := newEmployee(11, "Nombre Con Constructor")
	fmt.Println(*e4)
}
