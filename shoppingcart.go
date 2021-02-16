package shoppingcart

import "fmt"

type Cart map[string]int

func (c Cart) AddProduct(productCode string, quantity int) {
	if qty, ok := c[productCode]; ok {
		c[productCode] = qty + quantity
	} else {
		c[productCode] = quantity
	}
}

func (c Cart) RemoveProduct(productCode string) {
	delete(c, productCode)
}

func (c Cart) PrintCart() {
	for productCode, quantity := range c {
		fmt.Printf("%s (%d)\n", productCode, quantity)
	}
}
