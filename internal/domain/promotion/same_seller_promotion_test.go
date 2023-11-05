package promotion

import (
	"testing"
	"trendyol/internal/domain/item"
)

func TestSameSellerPromotionApply(t *testing.T) {
	testCases := []struct {
		items        []item.Item
		totalPrice   float64
		expectedID   int
		expectedDisc float64
	}{
		{
			items: []item.Item{
				item.NewItem(1, 2, 3, 4, 5.0),
				item.NewItem(6, 7, 3, 9, 10.0),
			},
			totalPrice:   15.0,
			expectedID:   SameSellerPromotionId,
			expectedDisc: 1.5,
		},
		{
			items: []item.Item{
				item.NewItem(1, 2, 3, 4, 5.0),
				item.NewItem(6, 7, 8, 9, 10.0),
			},
			totalPrice:   15.0,
			expectedID:   SameSellerPromotionId,
			expectedDisc: 0.0,
		},
	}

	for _, tc := range testCases {
		mockCart := &MockCart{
			items: tc.items,
			price: tc.totalPrice,
		}

		prom := NewSameSellerPromotion()
		appliedProm := prom.Apply(mockCart)

		if appliedProm.Id() != tc.expectedID {
			t.Errorf("Expected promotion id to be %d, got %d", tc.expectedID, appliedProm.Id())
		}

		if appliedProm.Discount() != tc.expectedDisc {
			t.Errorf("Expected discount to be %f, got %f", tc.expectedDisc, appliedProm.Discount())
		}
	}
}
