package abtraction_methods

import "fmt"

// Estructuras en lugar de clases

/*
Podemos realizar el proceso de abstracción a través de las estructuras.
Las estructuras son simplemente colecciones de campos
*/

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

// * utilizando el (c Course) como parámetro receptor automáticamente PrintClasses se convierte en un método de Course.

func (c *Course) ChangePrice(newPrice float64) {
	c.Price = newPrice
}

// PrintClasses Función que me permite imprimir las diferentes clases que contienen la estructura Course
func (c *Course) PrintClasses() { // Le indicamos que tiene un receptor "c" que será de tipo Course
	// ? El argumento receptor debe indicar claramente a qué hace referencia. En este caso lo hace a Course
	text := "Las clases son: "
	for _, class := range c.Classes {
		text += class + ", "
	}
	fmt.Println(text[:len(text)-2])
}

// ! Nota: Es mejor separar la lógica por temas de organización y orden
