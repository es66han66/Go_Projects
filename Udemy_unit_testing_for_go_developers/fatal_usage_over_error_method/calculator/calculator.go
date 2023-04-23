package calculator

import "errors"

type DiscountCalculator struct {
	minimumPurchaseAmount int
	discountAmount        int
}

func NewDiscountCalculator(minimumPurchaseAmount int, discountAmount int) (*DiscountCalculator, error) {

	if minimumPurchaseAmount <= 0 {
		return &DiscountCalculator{}, errors.New("coult not instantiate calculator with 0 or less minimum purchase amount")
	}

	return &DiscountCalculator{
		minimumPurchaseAmount: minimumPurchaseAmount,
		discountAmount:        discountAmount,
	}, nil
}

func (c *DiscountCalculator) Calculate(purchaseAmount int) int {
	if purchaseAmount > c.minimumPurchaseAmount {

		multiplier := purchaseAmount / c.minimumPurchaseAmount

		return purchaseAmount - (c.discountAmount * multiplier)
	}

	return purchaseAmount
}
