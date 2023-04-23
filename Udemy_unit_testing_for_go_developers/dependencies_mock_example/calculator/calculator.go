package calculator

import (
	"errors"
	"testing-example/database"
)

type DiscountCalculator struct {
	minimumPurchaseAmount int
	discountRepository    database.Repository
}

func NewDiscountCalculator(minimumPurchaseAmount int, discountRepository database.Repository) (*DiscountCalculator, error) {

	if minimumPurchaseAmount <= 0 {
		return &DiscountCalculator{}, errors.New("coult not instantiate calculator with 0 or less minimum purchase amount")
	}

	return &DiscountCalculator{
		minimumPurchaseAmount: minimumPurchaseAmount,
		discountRepository:    discountRepository,
	}, nil
}

func (c *DiscountCalculator) Calculate(purchaseAmount int) int {
	if purchaseAmount > c.minimumPurchaseAmount {

		multiplier := purchaseAmount / c.minimumPurchaseAmount

		discount := c.discountRepository.FindCurrentDiscount()

		return purchaseAmount - (discount * multiplier)
	}

	return purchaseAmount
}
