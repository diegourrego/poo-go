package main

import "fmt"

type Storager interface {
	Get() string
	Set(string)
}

type Person struct {
	name string
}

func NewPerson(name string) *Person {
	return &Person{name: name}
}

// Creamos los dos métodos para implementar la interface 'Storager'

func (p *Person) Get() string {
	return p.name
}

func (p *Person) Set(name string) {
	p.name = name
}

func Execute(s Storager, name string) {
	s.Set(name)
	fmt.Println(s.Get())
}

func main() {
	p := NewPerson("Jorge")
	Execute(p, "Ricardo")
}
