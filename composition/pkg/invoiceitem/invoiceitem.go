package invoiceitem

// Item contiene información de un item de factura
type Item struct {
	id      uint
	product string
	value   float64
}

func New(id uint, product string, value float64) Item {
	return Item{id, product, value}
}

// Value getter of Item.Value
func (i Item) Value() float64 {
	return i.value
}
