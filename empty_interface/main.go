package main

import (
	"fmt"
	"strings"
)

type exampler interface {
	x()
}

// Las interfaces nulas nos sirven para cuando queramos trabajar con tipos desconocidos
// usar any o interface{} es exactamente lo mismo.

func wrapper(i any) {
	//fmt.Printf("Valor: %v, Tipo: %T\n", i, i)

	// Type assertion
	// i == string? -> v: value, ok: es o no es
	//v, ok := i.(string)
	//if ok {
	//	fmt.Println(strings.ToUpper(v))
	//}

	// Switch type
	switch v := i.(type) {
	case string:
		fmt.Println(strings.ToUpper(v))
	case int:
		fmt.Println(v * 2)
	default:
		fmt.Printf("Valor: %v, Tipo: %T\n", i, i)
	}

}

func main() {
	//var e exampler // Interface nula
	//e.x()
	wrapper(12)
	wrapper("Jennyfer")
	wrapper(false)
	wrapper("Hola")
}
