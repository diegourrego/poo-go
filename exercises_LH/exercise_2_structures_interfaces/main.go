package main

import "fmt"

type Product interface {
	Price(productPrice float64) float64
}

// Productos

type Small struct {
}

func (s Small) Price(productPrice float64) float64 {
	return productPrice
}

type Medium struct {
}

func (m Medium) Price(productPrice float64) float64 {
	return productPrice + (productPrice * 0.03)
}

type Large struct {
}

func (l Large) Price(productPrice float64) float64 {
	var shippingCost = 2500.0
	return productPrice + (productPrice * 0.06) + shippingCost
}

func Factory(size uint) Product {
	switch size {
	case 1:
		return Small{}
	case 2:
		return Medium{}
	case 3:
		return Large{}
	default:
		return nil
	}
}

func main() {
	validSizeMessage := "You must enter a valid size"

	fmt.Println("Digita un número de tamaño de producto:")
	fmt.Println("\t 1: Small 2: Medium 3: Large")
	var size uint
	_, err := fmt.Scanln(&size)
	if err != nil {
		panic(validSizeMessage)
	}
	if size > 3 || size == 0 {
		panic(validSizeMessage)
	}
	product := Factory(size)
	fmt.Printf("The product has a price of: %v\n", product.Price(43500))
}
