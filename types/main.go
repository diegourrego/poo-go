package types

import "fmt"

type course struct {
	name string
}

func (c course) Print() {
	fmt.Printf("%+v\n", c)
}

// Declaraciones de alias
type myAlias = course // * Con esto podemos acceder a las propiedades y métodos de course

// Definición de tipo -→ Obtenemos el conjunto de campos, pero no heredamos sus métodos
type newCourse course // ! Podemos acceder únicamente a las propiedades

type newBool bool // * Al crear nuevos tipos pre declarados le podemos agregar métodos

func (b newBool) String() string {
	if b {
		return "VERDADERO"
	}
	return "FALSO"
}

func RunTypes() {
	//c := newCourse{name: "Go"}
	//c.Print()
	//fmt.Printf("El tipo es; %T\n", c)
	var b newBool = false
	fmt.Println(b.String())
}
