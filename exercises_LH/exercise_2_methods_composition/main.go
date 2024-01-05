package main

import (
	"fmt"
)

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

func NewPerson(id int, name, dateOfBirth string) Person {
	return Person{
		ID:          id,
		Name:        name,
		DateOfBirth: dateOfBirth,
	}
}

type Employee struct {
	Person
	ID       int
	Position string
}

func NewEmployee(id int, name, dateOfBirth string, position string) Employee {
	return Employee{
		NewPerson(id, name, dateOfBirth),
		id,
		position,
	}

}

func (e Employee) PrintEmployee(employee Employee) {
	fmt.Printf("Employee: %+v", employee)
}

func main() {
	// Instancias
	person := NewPerson(1, "Juan Hernandez", "1990-01-01")
	fmt.Printf("Person: %+v\n", person)
	employee := NewEmployee(2, "Eduardo Rojas", "1987-02-04", "Developer")

	employee.PrintEmployee(employee)
}
