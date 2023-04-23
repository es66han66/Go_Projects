package calculator

import "testing"

// func Test_DiscountCalculator(t *testing.T) {
// 	type testCase struct {
// 		minimumPurchaseAmount, discount, purchaseAmount, expectedAmount int
// 	}

// 	testCases := []testCase{
// 		{100, 20, 150, 130},
// 		{100, 20, 200, 160},
// 		{100, 20, 350, 290},
// 		{100, 20, 50, 60},
// 	}

// 	for _, tc := range testCases {
// 		c := NewDiscountCalculator(tc.minimumPurchaseAmount, tc.discount)
// 		amount := c.Calculate(tc.purchaseAmount)
// 		if amount != tc.expectedAmount {
// 			t.Errorf("Expected %d but got %d", tc.expectedAmount, amount)
// 		}
// 	}
// }

/*
In above scenario of writing test cases and running like above, which test case under the struct is failing and which are passing is hard to detect, below can be followed which is sub test approach
*/

func Test_DiscountCalculator1(t *testing.T) {
	type testCase struct {
		name                                                            string
		minimumPurchaseAmount, discount, purchaseAmount, expectedAmount int
	}

	testCases := []testCase{
		{"test1", 100, 20, 150, 130},
		{"test2", 100, 20, 200, 160},
		{"test3", 100, 20, 350, 290},
		{"test4", 100, 20, 50, 60},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := NewDiscountCalculator(tc.minimumPurchaseAmount, tc.discount)
			amount := c.Calculate(tc.purchaseAmount)
			if amount != tc.expectedAmount {
				t.Errorf("Expected %d but got %d", tc.expectedAmount, amount)
			}

		})
	}
}
