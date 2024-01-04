package main

import (
	"encapsulation/course"
	"fmt"
)

func main() {

	// Creamos el objeto
	Go := course.New("Go desde cero", 12.34, false) // Inicializamos la estructura
	Go.SetUserIDs([]uint{12, 56, 89})
	Go.SetClasses(map[uint]string{
		1: "Introducción",
		2: "Estructuras",
		3: "Maps",
	})

	Go.SetName("POO con Go")
	fmt.Println(Go.Name())
	Go.PrintClasses()
	//fmt.Printf("%+v", Go)
}
