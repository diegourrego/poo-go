package main

import (
	"errors"
	"fmt"
)

var products = []Product{
	{ID: 1, Name: "Macbook Pro", Price: 3200.34,
		Description: "14'' screen, M3 processor, 32GB RAM", Category: "Technology"},
	{ID: 2, Name: "Mouse Mx Master", Price: 520.62,
		Description: "Ultra velocity mouse", Category: "Technology"},
	{ID: 3, Name: "Beef", Price: 14.48,
		Description: "Cow meat", Category: "Food"},
	{ID: 4, Name: "T-Shirt The North Face", Price: 52.3,
		Description: "Comfortable t-shirt", Category: "Clothes"},
}

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

func (p Product) Save(product Product) {
	products = append(products, product)
}

func (p Product) GetAll() {
	fmt.Println("Products in cart:")
	for i, product := range products {
		fmt.Printf("%v. %v\n", i, product)
	}
}

func getById(id int) (Product, error) {
	for _, product := range products {
		if product.ID == id {
			return product, nil
		}
	}
	return Product{}, errors.New("product not found")
}

func main() {
	var newProduct Product = Product{ID: 5, Name: "iPhone 11", Price: 563.87, Description: "Cellphone", Category: "Technology"}
	fmt.Println("Cart before adding the new product")
	newProduct.GetAll()
	newProduct.Save(newProduct)
	fmt.Println("Cart after adding the new product")
	newProduct.GetAll()

	product, err := getById(1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("The product obtained is: %+v", product)
}
