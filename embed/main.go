package main

import "fmt"

type Person struct {
	Name string
	Age  uint
}

func NewPerson(name string, age uint) Person {
	return Person{
		Name: name,
		Age:  age,
	}
}

func (p Person) Greet() {
	fmt.Println("Hello")
}

type Human struct {
	Age      uint
	Children uint
}

func NewHuman(age, children uint) Human {
	return Human{age, children}
}

// La idea es poder embeber persona dentro de empleado

type Employee struct {
	// * Si solo especificamos el tipo de dato, automáticamente embebemos persona dentro de empleado.
	Person // Con esto le decimos a go que los campos y métodos se convertirán en los campos y métodos de empleado
	Human
	Salary float64
}

func NewEmployee(name string, age, children uint, salary float64) Employee {
	return Employee{
		NewPerson(name, age),
		NewHuman(age, children),
		salary,
	}
}

func (e Employee) Greet() {
	fmt.Println("Hello from employee ")
}

func (e Employee) Payroll() {
	fmt.Println(e.Salary * 0.90)
}

func main() {
	e := NewEmployee("Jennyfer", 27, 2, 1000000)
	fmt.Println(e.Name, e.Human.Age)
	e.Person.Greet()
	e.Payroll()
	e.Greet()
}
