package abtraction_methods

import "fmt"

func GoAbstraction() {

	// Creamos el objeto
	Go := Course{
		"Go desde Cero",
		12.34,
		false,
		[]uint{12, 56, 89},
		map[uint]string{
			1: "Introducción",
			2: "Estructuras",
			3: "Maps",
		},
	}

	css := Course{Name: "CSS desde cero", IsFree: true} // Al resto de campos go le asigna los valores cero
	js := Course{}
	js.Name = "Curso JavaScript"
	js.UserIDs = []uint{12, 67}
	js.Classes = map[uint]string{
		1: "Introducción y bienvenida",
		2: "let y const en JS",
		3: "Declaración de variables",
	}

	fmt.Println(Go.Name)
	fmt.Printf("%+v\n", css)
	fmt.Printf("%+v\n", js)

	Go.PrintClasses()
	//js.PrintClasses()
	Go.ChangePrice(67.12)
	fmt.Println(Go.Price)

}
