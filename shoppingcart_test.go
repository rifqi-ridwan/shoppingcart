package shoppingcart_test

import (
	"io/ioutil"
	"os"
	"shoppingcart"
	"testing"

	"github.com/go-playground/assert"
)

var (
	cart = &shoppingcart.Cart{
		"Mango": 2,
		"Apple": 3,
	}
)

func TestAddProduct(t *testing.T) {
	expected := &shoppingcart.Cart{
		"Mango": 2,
		"Apple": 3,
		"Melon": 4,
	}

	cart.AddProduct("Melon", 4)

	assert.Equal(t, expected, cart)
}

func TestRemoveProduct(t *testing.T) {
	err := cart.RemoveProduct("Mango")

	assert.Equal(t, nil, err)
}

func TestRemoveNotListedProduct(t *testing.T) {
	expected := "product is not listed"
	err := cart.RemoveProduct("Orange")

	assert.Equal(t, expected, err.Error())
}

func TestPrintCart(t *testing.T) {
	expected := "Mango (2)\nApple (3)\n"

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	cart.PrintCart()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	assert.Equal(t, expected, string(out))
}
