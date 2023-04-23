package calculator

import (
	"testing"
)

type DiscountRepositoryMock struct {
}

func (dc *DiscountRepositoryMock) FindCurrentDiscount() int {
	return 20
}

func Test_DiscountCalculator1(t *testing.T) {
	type testCase struct {
		name                                                            string
		minimumPurchaseAmount, discount, purchaseAmount, expectedAmount int
	}

	testCases := []testCase{
		{"test1", 100, 20, 150, 130},
		{"test2", 100, 20, 200, 160},
		{"test3", 100, 20, 350, 290},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			discountRepoMock := &DiscountRepositoryMock{}

			c, err := NewDiscountCalculator(tc.minimumPurchaseAmount, discountRepoMock)

			if err != nil {
				t.Fatal(err.Error())
			}

			amount := c.Calculate(tc.purchaseAmount)
			if amount != tc.expectedAmount {
				t.Fatalf("Expected %d but got %d", tc.expectedAmount, amount)
			}

		})
	}
}
