package course

import "fmt"

// Estructuras en lugar de clases

/*
Podemos realizar el proceso de abstracción a través de las estructuras.
Las estructuras son simplemente colecciones de campos
*/

// Un identificador exportado se identifica por tener la primera letra en mayúscula

type course struct {
	name    string
	price   float64
	isFree  bool // identificador no exportado
	userIDs []uint
	classes map[uint]string
}

// * Métodos de acceso - Setter y Getters

func (c *course) SetName(name string) {
	c.name = name
}

// Name método getter. En Go por buenas prácticas la palabra "get" se debe omitir en el nombre de la función
func (c *course) Name() string {
	return c.name
}

func (c *course) SetPrice(price float64) {
	c.price = price
}

func (c *course) Price() float64 {
	return c.price
}

func (c *course) SetIsFree(isFree bool) {
	c.isFree = isFree
}

func (c *course) IsFree() bool {
	return c.isFree
}

func (c *course) SetUserIDs(usersIDs []uint) {
	c.userIDs = usersIDs
}

func (c *course) UserIDs() []uint {
	return c.userIDs
}

func (c *course) SetClasses(classes map[uint]string) {
	c.classes = classes
}

func (c *course) Classes() map[uint]string {
	return c.classes
}

// New Función constructora
func New(name string, price float64, isFree bool) *course {

	if price == 0 {
		price = 30
	}

	return &course{
		name:   name,
		price:  price,
		isFree: isFree,
	}
}

// PrintClasses Función que me permite imprimir las diferentes clases que contienen la estructura course
func (c *course) PrintClasses() { // Le indicamos que tiene un receptor "c" que será de tipo Course
	// ? El argumento receptor debe indicar claramente a qué hace referencia. En este caso lo hace a Course
	text := "Las clases son: "
	for _, class := range c.classes {
		text += class + ", "
	}
	fmt.Println(text[:len(text)-2])
}

// ! Nota: Es mejor separar la lógica por temas de organización y orden
