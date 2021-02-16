package shoppingcart_test

import (
	"io/ioutil"
	"os"
	"shoppingcart"
	"testing"

	"github.com/go-playground/assert"
)

func TestFullScenario(t *testing.T) {
	expected := "Green Banana (2)\nRed Apple (7)\n"

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	cart := shoppingcart.Cart{}

	cart.AddProduct("Green Banana", 2)

	cart.AddProduct("Yellow Watermelon", 3)

	cart.AddProduct("Red Apple", 1)
	cart.AddProduct("Red Apple", 4)
	cart.AddProduct("Red Apple", 2)

	cart.RemoveProduct("Yellow Watermelon")

	cart.RemoveProduct("Red Watermelon")

	cart.PrintCart()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	assert.Equal(t, expected, string(out))
}

func TestAddProduct(t *testing.T) {
	expected := &shoppingcart.Cart{
		"Mango": 2,
		"Apple": 3,
	}

	cart := shoppingcart.Cart{}

	cart.AddProduct("Mango", 2)
	cart.AddProduct("Apple", 3)

	assert.Equal(t, expected, cart)
}

func TestRemoveProduct(t *testing.T) {
	expected := &shoppingcart.Cart{
		"Mango": 2,
	}

	cart := shoppingcart.Cart{}

	cart.AddProduct("Mango", 2)
	cart.AddProduct("Apple", 3)

	cart.RemoveProduct("Apple")

	assert.Equal(t, expected, cart)
}

func TestPrintCart(t *testing.T) {
	expected := "Mango (2)\nApple (3)\n"

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	cart := shoppingcart.Cart{}

	cart.AddProduct("Mango", 2)
	cart.AddProduct("Apple", 3)

	cart.PrintCart()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	assert.Equal(t, expected, string(out))
}
