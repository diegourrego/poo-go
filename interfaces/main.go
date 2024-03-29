package main

import "fmt"

type Greeter interface {
	Greet()
}

type Byer interface {
	Bye()
}

// GreeterByer interfaz embebida
type GreeterByer interface {
	Greeter
	Byer
}

type Person struct {
	Name string
}

func (p Person) Greet() {
	fmt.Printf("Hola soy %s\n", p.Name)
}

func (p Person) Bye() {
	fmt.Printf("Adiós soy %s\n", p.Name)
}

func (p Person) String() string {
	return "Hola soy una persona y mi nombre es: " + p.Name
}

type Text string

func (t Text) Greet() {
	fmt.Printf("Hola soy %s\n", t)
}

func (t Text) Bye() {
	fmt.Printf("Adiós soy %s\n", t)
}

func All(gbs ...GreeterByer) {
	for _, gb := range gbs {
		gb.Greet()
		gb.Bye()
	}
}

func main() {
	var p = Person{Name: "Diego"}
	//var t = Text("Jennyfer")
	//All(p, t)
	fmt.Println(p)
}
