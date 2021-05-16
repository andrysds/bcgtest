package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitInventory(t *testing.T) {
	item := &InventoryItem{
		SKU:   "120P90",
		Name:  "Google Home",
		Price: 49.99,
		Qty:   10,
	}
	items := []*InventoryItem{item}
	expected := map[string]*InventoryItem{item.SKU: item}

	InitInventory(items)
	assert.Equal(t, expected, Inventory)
}
