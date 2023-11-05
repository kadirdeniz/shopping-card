package item

import (
	"fmt"
	"testing"
)

func TestVasItemValidation(t *testing.T) {
	testCases := []struct {
		categoryId    int
		sellerId      int
		expectedError error
	}{
		{VasItemCategoryId, VasItemSellerId, nil},
		{VasItemCategoryId, 1234, ErrInvalidSellerId},
		{1234, VasItemSellerId, ErrInvalidCategoryId},
		{1234, 1234, ErrInvalidCategoryId}, // SellerId won't be checked if CategoryId is invalid
	}

	for _, tc := range testCases {
		item := vasItem{item{id: 1, categoryId: tc.categoryId, sellerId: tc.sellerId, price: 1.0, quantity: 1}}
		err := item.Validate()
		if err != tc.expectedError {
			t.Error(fmt.Sprintf("Expected error %v, got %v", tc.expectedError, err))
		}
	}
}
