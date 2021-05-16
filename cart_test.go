package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCartAddItem(t *testing.T) {
	item := &InventoryItem{
		SKU:   "120P90",
		Name:  "Google Home",
		Price: 49.99,
		Qty:   10,
	}
	items := []*InventoryItem{item}
	InitInventory(items)

	cases := []struct {
		cart                *Cart
		sku                 string
		qty                 int
		isError             bool
		expectedCartItems   map[string]*CartItem
		expectedTotalAmount float64
	}{
		// adding not found item
		{
			cart: &Cart{
				Items:       map[string]*CartItem{},
				TotalAmount: 0,
			},
			sku:                 "234324",
			qty:                 1,
			isError:             true,
			expectedCartItems:   map[string]*CartItem{},
			expectedTotalAmount: 0,
		},
		// qty is greater than inventory qty
		{
			cart: &Cart{
				Items:       map[string]*CartItem{},
				TotalAmount: 0,
			},
			sku:                 "120P90",
			qty:                 100,
			isError:             true,
			expectedCartItems:   map[string]*CartItem{},
			expectedTotalAmount: 0,
		},
		// adding new cart item
		{
			cart: &Cart{
				Items:       map[string]*CartItem{},
				TotalAmount: 0,
			},
			sku:     "120P90",
			qty:     1,
			isError: false,
			expectedCartItems: map[string]*CartItem{
				"120P90": {
					SKU:    "120P90",
					Qty:    1,
					Amount: item.Price,
				},
			},
			expectedTotalAmount: item.Price,
		},
		// adding more qty to existing cart item
		{
			cart: &Cart{
				Items: map[string]*CartItem{
					"120P90": {
						SKU:    "120P90",
						Qty:    1,
						Amount: item.Price,
					},
				},
				TotalAmount: item.Price,
			},
			sku:     "120P90",
			qty:     1,
			isError: false,
			expectedCartItems: map[string]*CartItem{
				"120P90": {
					SKU:    "120P90",
					Qty:    2,
					Amount: 2 * item.Price,
				},
			},
			expectedTotalAmount: 2 * item.Price,
		},
	}

	for _, c := range cases {
		err := c.cart.AddItem(c.sku, c.qty)

		if c.isError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, c.expectedCartItems, c.cart.Items)
			assert.Equal(t, c.expectedTotalAmount, c.cart.TotalAmount)
		}
	}
}
