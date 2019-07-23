package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBasketOK(t *testing.T) {
	var e Basket
	b := CreateBasket()
	assert.IsType(t, e, b)
}

func TestAddProductOK(t *testing.T) {
	CreateBasket()
	err := AddProduct(1, "VOUCHER")
	assert.Nil(t, err)
}

func TestAddProductError(t *testing.T) {
	err := AddProduct(10, "VOUCHER")
	assert.Error(t, err)
}

func TestGetAmountOK(t *testing.T) {
	CreateBasket()
	err := AddProduct(1, "VOUCHER")
	assert.Nil(t, err)
	r, err := GetAmount(1)
	e := 10.00
	var typeOf float64
	assert.IsType(t, typeOf, r)
	assert.Equal(t, e, r)
}

func TestGetAmountError(t *testing.T) {
	CreateBasket()
	_, err := GetAmount(1000)
	assert.Error(t, err)
}

func TestRemoveBasketOK(t *testing.T) {
	CreateBasket()
	err := RemoveBasket(1)
	assert.Nil(t, err)
}

func TestRemoveBasketError(t *testing.T) {
	CreateBasket()
	err := RemoveBasket(100)
	assert.Error(t, err)
}
