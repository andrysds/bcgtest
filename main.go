//+build !test

package main

import "fmt"

func main() {
	items := []*InventoryItem{
		{
			SKU:   "120P90",
			Name:  "Google Home",
			Price: 49.99,
			Qty:   10,
		},
		{
			SKU:   "43N23P",
			Name:  "MacBook Pro",
			Price: 5399.99,
			Qty:   5,
		},
		{
			SKU:   "A304SD",
			Name:  "Alexa Speaker",
			Price: 109.5,
			Qty:   10,
		},
		{
			SKU:   "234234",
			Name:  "Raspberry Pi B",
			Price: 30,
			Qty:   2,
		},
	}
	InitInventory(items)

	fmt.Println(Inventory)
}
