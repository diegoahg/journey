package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddItemOK(t *testing.T) {
	var i Item
	i.FillItem("VOUCHER")
	var b Basket
	b.AddItem(i)
	i.FillItem("TSHIRT")
	b.AddItem(i)
	i.FillItem("TSHIRT")
	b.AddItem(i)
	i.FillItem("TSHIRT")
	b.AddItem(i)
	i.FillItem("TSHIRT")
	b.AddItem(i)
	i.FillItem("MUG")
	b.AddItem(i)
	r := b.GetTotal()
	e := 88.50
	var typeOf float64
	assert.IsType(t, typeOf, r)
	assert.Equal(t, e, r)
}
