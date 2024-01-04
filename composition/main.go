package main

import (
	"composition/pkg/customer"
	"composition/pkg/invoice"
	"composition/pkg/invoiceitem"
	"fmt"
)

func main() {

	// Ejemplo composición
	jorge := customer.New("Jorge Español", "Av Calle 123 #45-67", "3101234567")

	var jorgeItems = invoiceitem.NewItems(
		invoiceitem.New(1, "Tomates", 2500),
		invoiceitem.New(2, "Salsa de ajo", 8750),
		invoiceitem.New(3, "Carne de res", 25300),
		invoiceitem.New(4, "Paca de leche", 38000),
	)

	invoice1 := invoice.New("Colombia", "Bogotá D.C", jorge, jorgeItems)
	invoice1.SetTotal()
	fmt.Printf("%+v", invoice1)
}
