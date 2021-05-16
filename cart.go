package main

import "errors"

type CartItem struct {
	SKU    string
	Qty    int
	Amount float64
}

type Cart struct {
	Items       map[string]*CartItem
	TotalAmount float64
}

func (c *Cart) AddItem(sku string, qty int) error {
	item, err := GetInventoryItem(sku)
	if err != nil {
		return err
	}

	if qty > item.Qty {
		return errors.New("qty requested is not available")
	}

	amount := float64(qty) * item.Price
	cartItem, exists := c.Items[sku]
	if exists {
		cartItem.Qty += qty
		cartItem.Amount += amount
	} else {
		newCartItem := &CartItem{
			SKU:    sku,
			Qty:    qty,
			Amount: amount,
		}
		c.Items[sku] = newCartItem
	}

	item.Qty -= qty
	c.TotalAmount += amount

	return nil
}
