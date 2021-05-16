package main

type InventoryItem struct {
	SKU   string
	Name  string
	Price float64
	Qty   int
}

var Inventory map[string]*InventoryItem

func InitInventory(items []*InventoryItem) {
	Inventory = map[string]*InventoryItem{}

	for _, item := range items {
		Inventory[item.SKU] = item
	}
}
