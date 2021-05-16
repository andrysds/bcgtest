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
	}
	InitInventory(items)

	fmt.Println(Inventory)
}
