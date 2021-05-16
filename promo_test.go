package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFreeItemPromoCalculateDiscount(t *testing.T) {
	promo := &FreeItemPromo{
		Description: "Buy MacBook Pro, Free Raspberry Pi B",
		ReqItemSKU:  "43N23P",
		FreeItemSKU: "234234",
	}

	cases := []struct {
		cart             *Cart
		expectedDiscount *Discount
		isError          bool
	}{
		// buy 1 MacBook Pro & 1 Raspberry Pi B
		{
			cart: &Cart{
				Items: map[string]*CartItem{
					"43N23P": {
						SKU:    "43N23P",
						Qty:    1,
						Amount: 5399.99,
					},
					"234234": {
						SKU:    "234234",
						Qty:    1,
						Amount: 30,
					},
				},
				TotalAmount: 5399.99 + 30,
			},
			expectedDiscount: &Discount{
				Description: promo.Description,
				Amount:      30,
			},
			isError: false,
		},
		// buy 2 MacBook Pro & 1 Raspberry Pi B
		{
			cart: &Cart{
				Items: map[string]*CartItem{
					"43N23P": {
						SKU:    "43N23P",
						Qty:    2,
						Amount: 2 * 5399.99,
					},
					"234234": {
						SKU:    "234234",
						Qty:    1,
						Amount: 30,
					},
				},
				TotalAmount: 2*5399.99 + 30,
			},
			expectedDiscount: &Discount{
				Description: promo.Description,
				Amount:      30,
			},
			isError: false,
		},
		// buy 1 MacBook Pro & 2 Raspberry Pi B
		{
			cart: &Cart{
				Items: map[string]*CartItem{
					"43N23P": {
						SKU:    "43N23P",
						Qty:    1,
						Amount: 5399.99,
					},
					"234234": {
						SKU:    "234234",
						Qty:    2,
						Amount: 2 * 30,
					},
				},
				TotalAmount: 5399.99 + 2*30,
			},
			expectedDiscount: &Discount{
				Description: promo.Description,
				Amount:      30,
			},
			isError: false,
		},
		// buy 1 MacBook Pro only
		{
			cart: &Cart{
				Items: map[string]*CartItem{
					"43N23P": {
						SKU:    "43N23P",
						Qty:    1,
						Amount: 5399.99,
					},
				},
				TotalAmount: 5399.99,
			},
			expectedDiscount: nil,
			isError:          true,
		},
		// buy 1 Raspberry Pi B only
		{
			cart: &Cart{
				Items: map[string]*CartItem{
					"234234": {
						SKU:    "234234",
						Qty:    1,
						Amount: 30,
					},
				},
				TotalAmount: 5399.99,
			},
			expectedDiscount: nil,
			isError:          true,
		},
	}

	for _, c := range cases {
		actualDiscount, err := promo.CalculateDiscount(c.cart)

		assert.Equal(t, c.expectedDiscount, actualDiscount)
		if c.isError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}

	}
}

func TestFreeSameItemPromoCalculateDiscount(t *testing.T) {
	promo := &FreeSameItemPromo{
		Description: "Buy 3 Google Homes for the price of 2",
		ReqItemSKU:  "120P90",
		ReqItemNum:  3,
		FreeItemNum: 1,
	}

	cases := []struct {
		cart             *Cart
		expectedDiscount *Discount
		isError          bool
	}{
		// buy 3 Google Home
		{
			cart: &Cart{
				Items: map[string]*CartItem{
					"120P90": {
						SKU:    "120P90",
						Qty:    3,
						Amount: 3 * 49.99,
					},
				},
				TotalAmount: 3 * 49.99,
			},
			expectedDiscount: &Discount{
				Description: promo.Description,
				Amount:      49.99,
			},
			isError: false,
		},
		// buy 7 Google Home
		{
			cart: &Cart{
				Items: map[string]*CartItem{
					"120P90": {
						SKU:    "120P90",
						Qty:    7,
						Amount: 7 * 49.99,
					},
				},
				TotalAmount: 7 * 49.99,
			},
			expectedDiscount: &Discount{
				Description: promo.Description,
				Amount:      2 * 49.99,
			},
			isError: false,
		},
		// buy 1 Google Home
		{
			cart: &Cart{
				Items: map[string]*CartItem{
					"120P90": {
						SKU:    "120P90",
						Qty:    1,
						Amount: 49.99,
					},
				},
				TotalAmount: 49.99,
			},
			expectedDiscount: nil,
			isError:          true,
		},
		// buy 1 Alexa Speaker
		{
			cart: &Cart{
				Items: map[string]*CartItem{
					"A304SD": {
						SKU:    "A304SD",
						Qty:    1,
						Amount: 109.5,
					},
				},
				TotalAmount: 109.5,
			},
			expectedDiscount: nil,
			isError:          true,
		},
	}

	for _, c := range cases {
		actualDiscount, err := promo.CalculateDiscount(c.cart)

		assert.Equal(t, c.expectedDiscount, actualDiscount)
		if c.isError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}

	}
}

func TestPercentageDiscountPromoCalculateDiscount(t *testing.T) {
	promo := &PercentageDiscountPromo{
		Description:         "Buy 3 or more Alexa Speakers, 10% off",
		ReqItemSKU:          "A304SD",
		ReqItemMoreEqualNum: 3,
		DiscountPercentage:  0.1,
	}

	cases := []struct {
		cart             *Cart
		expectedDiscount *Discount
		isError          bool
	}{
		// buy 3 Alexa Speakers
		{
			cart: &Cart{
				Items: map[string]*CartItem{
					"A304SD": {
						SKU:    "A304SD",
						Qty:    3,
						Amount: 3 * 109.5,
					},
				},
				TotalAmount: 3 * 109.5,
			},
			expectedDiscount: &Discount{
				Description: promo.Description,
				Amount:      0.1 * 3 * 109.5,
			},
			isError: false,
		},
		// buy 7 Alexa Speakers
		{
			cart: &Cart{
				Items: map[string]*CartItem{
					"A304SD": {
						SKU:    "A304SD",
						Qty:    7,
						Amount: 7 * 109.5,
					},
				},
				TotalAmount: 7 * 109.5,
			},
			expectedDiscount: &Discount{
				Description: promo.Description,
				Amount:      0.1 * 7 * 109.5,
			},
			isError: false,
		},
		// buy 1 Alexa Speaker
		{
			cart: &Cart{
				Items: map[string]*CartItem{
					"A304SD": {
						SKU:    "A304SD",
						Qty:    1,
						Amount: 109.5,
					},
				},
				TotalAmount: 109.5,
			},
			expectedDiscount: nil,
			isError:          true,
		},
		// buy 1 Google Home
		{
			cart: &Cart{
				Items: map[string]*CartItem{
					"120P90": {
						SKU:    "120P90",
						Qty:    1,
						Amount: 49.99,
					},
				},
				TotalAmount: 49.99,
			},
			expectedDiscount: nil,
			isError:          true,
		},
	}

	for _, c := range cases {
		actualDiscount, err := promo.CalculateDiscount(c.cart)

		assert.Equal(t, c.expectedDiscount, actualDiscount)
		if c.isError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}

	}
}
