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

func TestGetInventoryItem(t *testing.T) {
	item := &InventoryItem{
		SKU:   "120P90",
		Name:  "Google Home",
		Price: 49.99,
		Qty:   10,
	}
	items := []*InventoryItem{item}
	InitInventory(items)

	cases := []struct {
		sku          string
		expectedItem *InventoryItem
		isError      bool
	}{
		{
			sku:          "120P90",
			expectedItem: item,
			isError:      false,
		},
		{
			sku:          "234234",
			expectedItem: nil,
			isError:      true,
		},
	}

	for _, c := range cases {
		actualItem, err := GetInventoryItem(c.sku)
		assert.Equal(t, c.expectedItem, actualItem)
		if c.isError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}
