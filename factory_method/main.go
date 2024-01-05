package main

import "fmt"

type PayMethod interface {
	Pay()
}

type Paypal struct {
}

func (p Paypal) Pay() {
	fmt.Println("Pay it using Paypal")
}

type Cash struct {
}

func (c Cash) Pay() {
	fmt.Println("Pay it using Cash")

}

type CreditCard struct {
}

func (c CreditCard) Pay() {
	fmt.Println("Pay it using Credit card")

}

func Factory(method uint) PayMethod {
	switch method {
	case 1:
		return Paypal{}
	case 2:
		return Cash{}
	case 3:
		return CreditCard{}
	default:
		return nil
	}
}

func main() {
	var method uint

	fmt.Println("Digite un número de método de pago:")
	fmt.Println("\t 1: Paypal 2: Efectivo 3: Tarjeta de Crédito")
	_, err := fmt.Scanln(&method)
	if err != nil {
		panic("Debe digitar un método válido")
	}

	if method > 3 {
		panic("Debe digitar un método válido")
	}

	payMethod := Factory(method)
	payMethod.Pay()
}
