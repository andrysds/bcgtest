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

	promo := &FreeSameItemPromo{
		ID:          2,
		Description: "Buy 3 Google Homes for the price of 2",
		ReqItemSKU:  item.SKU,
		ReqItemNum:  3,
		FreeItemNum: 1,
	}
	InitPromos(map[string]Promo{item.SKU: promo})

	cases := []struct {
		cart                *Cart
		sku                 string
		qty                 int
		isError             bool
		expectedCartItems   map[string]*CartItem
		expectedDiscounts   map[int]*Discount
		expectedTotalAmount float64
	}{
		// adding not found item
		{
			cart: &Cart{
				Items:       map[string]*CartItem{},
				Discounts:   map[int]*Discount{},
				TotalAmount: 0,
			},
			sku:                 "234324",
			qty:                 1,
			isError:             true,
			expectedCartItems:   map[string]*CartItem{},
			expectedDiscounts:   map[int]*Discount{},
			expectedTotalAmount: 0,
		},
		// qty is greater than inventory qty
		{
			cart: &Cart{
				Items:       map[string]*CartItem{},
				Discounts:   map[int]*Discount{},
				TotalAmount: 0,
			},
			sku:                 item.SKU,
			qty:                 100,
			isError:             true,
			expectedCartItems:   map[string]*CartItem{},
			expectedDiscounts:   map[int]*Discount{},
			expectedTotalAmount: 0,
		},
		// adding new cart item
		{
			cart: &Cart{
				Items:       map[string]*CartItem{},
				Discounts:   map[int]*Discount{},
				TotalAmount: 0,
			},
			sku:     item.SKU,
			qty:     1,
			isError: false,
			expectedCartItems: map[string]*CartItem{
				item.SKU: {
					SKU:    item.SKU,
					Qty:    1,
					Amount: item.Price,
				},
			},
			expectedDiscounts:   map[int]*Discount{},
			expectedTotalAmount: item.Price,
		},
		// adding more qty to existing cart item
		{
			cart: &Cart{
				Items: map[string]*CartItem{
					item.SKU: {
						SKU:    item.SKU,
						Qty:    1,
						Amount: item.Price,
					},
				},
				Discounts:   map[int]*Discount{},
				TotalAmount: item.Price,
			},
			sku:     item.SKU,
			qty:     1,
			isError: false,
			expectedCartItems: map[string]*CartItem{
				item.SKU: {
					SKU:    item.SKU,
					Qty:    2,
					Amount: 2 * item.Price,
				},
			},
			expectedDiscounts:   map[int]*Discount{},
			expectedTotalAmount: 2 * item.Price,
		},
		// adding cart items and got promo
		{
			cart: &Cart{
				Items:       map[string]*CartItem{},
				Discounts:   map[int]*Discount{},
				TotalAmount: 0,
			},
			sku:     item.SKU,
			qty:     3,
			isError: false,
			expectedCartItems: map[string]*CartItem{
				item.SKU: {
					SKU:    item.SKU,
					Qty:    3,
					Amount: 3 * item.Price,
				},
			},
			expectedDiscounts: map[int]*Discount{
				1: {
					PromoID:     promo.ID,
					Description: promo.Description,
					Amount:      item.Price,
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
