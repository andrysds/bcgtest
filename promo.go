package main

import "errors"

var (
	ErrFreeItemNotFound = errors.New("free item not found")
	ErrReqItemNotFound  = errors.New("required item not found")
	ErrReqNotFulfilled  = errors.New("requirement not fulfilled")
)

type Discount struct {
	Description string
	Amount      float64
}

type Promo interface {
	CalculateDiscount()
}

type FreeItemPromo struct {
	Description string
	ReqItemSKU  string
	FreeItemSKU string
}

func (p *FreeItemPromo) CalculateDiscount(cart *Cart) (*Discount, error) {
	freeItem, found := cart.Items[p.FreeItemSKU]
	if !found {
		return nil, ErrFreeItemNotFound
	}

	reqItem, found := cart.Items[p.ReqItemSKU]
	if !found {
		return nil, ErrReqItemNotFound
	}

	freeNum := freeItem.Qty
	if reqItem.Qty < freeItem.Qty {
		freeNum = reqItem.Qty
	}

	unitPrice := freeItem.Amount / float64(freeItem.Qty)
	discount := &Discount{
		Description: p.Description,
		Amount:      float64(freeNum) * unitPrice,
	}

	return discount, nil
}

type FreeSameItemPromo struct {
	Description string
	ReqItemSKU  string
	ReqItemNum  int
	FreeItemNum int
}

func (p *FreeSameItemPromo) CalculateDiscount(cart *Cart) (*Discount, error) {
	item, found := cart.Items[p.ReqItemSKU]
	if !found {
		return nil, ErrReqItemNotFound
	}

	if item.Qty < p.ReqItemNum {
		return nil, ErrReqNotFulfilled
	}

	freeNum := item.Qty / p.ReqItemNum * p.FreeItemNum
	unitPrice := item.Amount / float64(item.Qty)
	discount := &Discount{
		Description: p.Description,
		Amount:      float64(freeNum) * unitPrice,
	}

	return discount, nil
}

type PercentageDiscountPromo struct {
	Description         string
	ReqItemSKU          string
	ReqItemMoreEqualNum int
	DiscountPercentage  float64
}

func (p *PercentageDiscountPromo) CalculateDiscount(cart *Cart) (*Discount, error) {
	item, found := cart.Items[p.ReqItemSKU]
	if !found {
		return nil, ErrReqItemNotFound
	}

	if item.Qty < p.ReqItemMoreEqualNum {
		return nil, ErrReqNotFulfilled
	}

	discount := &Discount{
		Description: p.Description,
		Amount:      item.Amount * p.DiscountPercentage,
	}

	return discount, nil
}
