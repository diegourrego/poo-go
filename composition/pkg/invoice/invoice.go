package invoice

import (
	"composition/pkg/customer"
	"composition/pkg/invoiceitem"
)

// Composición de la factura

type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer // * De esta manera hacemos la composición de 1 a 1
	items   invoiceitem.Items // * Relación de 1 a muchos
}

// New retorna una nueva factura
func New(country, city string, client customer.Customer, items invoiceitem.Items) Invoice {
	return Invoice{
		country: country,
		city:    city,
		client:  client,
		items:   items,
	}
}

// SetTotal Método que nos permita establecer el total
func (i *Invoice) SetTotal() {
	i.total += i.items.Total()
}
