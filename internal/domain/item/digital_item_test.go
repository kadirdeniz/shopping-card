package item

import (
	"testing"
)

func TestDigitalItemValidate(t *testing.T) {
	testCases := []struct {
		id         int
		categoryId int
		sellerId   int
		quantity   int
		price      float64
		totalPrice float64
		validate   error
	}{
		{1, DigitalItemCategoryId, 3, 4, 5.0, 20.0, nil},
		{2, DigitalItemCategoryId, 3, 6, 5.0, 30.0, ErrMaxQuantityForDigitalItem},
		{3, DigitalItemCategoryId, 3, 2, 5.0, 10.0, nil},
	}

	for _, tc := range testCases {
		digitalItem := digitalItem{item{tc.id, tc.categoryId, tc.sellerId, tc.price, tc.quantity}}

		// Test Validate()
		if err := digitalItem.Validate(); err != tc.validate {
			t.Errorf("Expected Validate to return %v, got %v", tc.validate, err)
		}
	}
}
