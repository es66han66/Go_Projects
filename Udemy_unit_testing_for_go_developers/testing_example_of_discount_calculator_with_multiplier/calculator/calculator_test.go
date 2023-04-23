package calculator

import "testing"

func Test_DiscountCalculator(t *testing.T) {
	type testCase struct {
		minimumPurchaseAmount, discount, purchaseAmount, expectedAmount int
	}

	testCases := []testCase{
		{100, 20, 150, 130},
		{100, 20, 200, 160},
		{100, 20, 350, 290},
		{100, 20, 50, 50},
	}

	for _, tc := range testCases {
		c := NewDiscountCalculator(tc.minimumPurchaseAmount, tc.discount)
		amount := c.Calculate(tc.purchaseAmount)
		if amount != tc.expectedAmount {
			t.Errorf("Expected %d but got %d", tc.expectedAmount, amount)
		}
	}
}
