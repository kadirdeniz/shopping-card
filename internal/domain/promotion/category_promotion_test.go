package promotion

import (
	"testing"
	"trendyol/internal/domain/item"
)

func TestCategoryPromotionApply(t *testing.T) {
	testCases := []struct {
		items        []item.Item
		expectedID   int
		expectedDisc float64
	}{
		{
			items: []item.Item{
				item.NewItem(1, DiscountedCategoryId, 3, 1, 100.0),
				item.NewItem(6, DiscountedCategoryId, 8, 1, 100.0),
				item.NewItem(11, DiscountedCategoryId, 13, 1, 100.0),
				item.NewItem(16, 17, 18, 19, 20.0),
			},
			expectedID:   CategoryPromotionId,
			expectedDisc: 15.0,
		},
		{
			items: []item.Item{
				item.NewItem(1, DiscountedCategoryId, 3, 2, 100.0),
				item.NewItem(6, DiscountedCategoryId, 8, 2, 100.0),
				item.NewItem(11, DiscountedCategoryId, 13, 2, 100.0),
				item.NewItem(16, 17, 18, 19, 20.0),
			},
			expectedID:   CategoryPromotionId,
			expectedDisc: 30.0,
		},
		{

			items: []item.Item{
				item.NewItem(1, 1, 3, 4, 5.0),
				item.NewItem(6, 1, 8, 9, 10.0),
				item.NewItem(11, 1, 13, 14, 15.0),
				item.NewItem(16, 1, 18, 19, 20.0),
			},
			expectedID:   CategoryPromotionId,
			expectedDisc: 0.0,
		},
	}

	for _, tc := range testCases {
		mockCart := &MockCart{
			items: tc.items,
		}

		catProm := NewCategoryPromotion()
		appliedProm := catProm.Apply(mockCart)

		if appliedProm.Id() != tc.expectedID {
			t.Errorf("Expected promotion id to be %d, got %d", tc.expectedID, appliedProm.Id())
		}

		if appliedProm.Discount() != tc.expectedDisc {
			t.Errorf("Expected discount to be %f, got %f", tc.expectedDisc, appliedProm.Discount())
		}
	}
}
