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

	promos := map[string]Promo{
		"43N23P": &FreeItemPromo{
			ID:          1,
			Description: "Buy MacBook Pro, Free Raspberry Pi B",
			ReqItemSKU:  "43N23P",
			FreeItemSKU: "234234",
		},
		"234234": &FreeItemPromo{
			ID:          1,
			Description: "Buy MacBook Pro, Free Raspberry Pi B",
			ReqItemSKU:  "43N23P",
			FreeItemSKU: "234234",
		},
		"120P90": &FreeSameItemPromo{
			ID:          2,
			Description: "Buy 3 Google Homes for the price of 2",
			ReqItemSKU:  "120P90",
			ReqItemNum:  3,
			FreeItemNum: 1,
		},
		"A304SD": &PercentageDiscountPromo{
			ID:                  3,
			Description:         "Buy 3 or more Alexa Speakers, 10% off",
			ReqItemSKU:          "A304SD",
			ReqItemMoreEqualNum: 3,
			DiscountPercentage:  0.1,
		},
	}
	InitPromos(promos)

	fmt.Println("Example Scenarios:")

	cart1 := NewCart()
	cart1.AddItem("43N23P", 1)
	cart1.AddItem("234234", 1)
	fmt.Println("Scanned Items: MacBook Pro, Raspberry Pi B")
	fmt.Printf("Total: $%.2f\n", cart1.TotalAmount) // Total: $5,399.99

	cart2 := NewCart()
	cart2.AddItem("120P90", 3)
	fmt.Println("Scanned Items: Google Home, Google Home, Google Home")
	fmt.Printf("Total: $%.2f\n", cart2.TotalAmount) // Total: $$99.98

	cart3 := NewCart()
	cart3.AddItem("A304SD", 3)
	fmt.Println("Scanned Items: Alexa Speaker, Alexa Speaker, Alexa Speaker")
	fmt.Printf("Total: $%.2f\n", cart3.TotalAmount) // Total: $295.65

}
