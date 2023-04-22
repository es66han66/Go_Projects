package calculator

import "testing"

func Test_DiscountApplied(t *testing.T) {

	c := NewDiscountCalculator(100, 20)

	amount := c.Calculate(120)

	if amount != 100 {
		t.Logf("Expected %d but got %d", 100, amount)
		t.Fail()
	}

	t.Log("hello test passed")

}

func Test_DiscountNotApplied(t *testing.T) {

	c := NewDiscountCalculator(100, 20)

	amount := c.Calculate(60)

	if amount != 60 {
		// t.Logf("Expected %d but got %d", 60, amount)
		// t.Fail()

		// in place of above below can be used

		t.Errorf("Expected %d but got %d", 60, amount)
	}

}
