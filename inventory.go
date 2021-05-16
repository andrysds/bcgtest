package main

import "errors"

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

func GetInventoryItem(sku string) (*InventoryItem, error) {
	item, found := Inventory[sku]
	if !found {
		return nil, errors.New("item not found")
	}
	return item, nil
}
