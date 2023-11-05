package service

import (
	"testing"
	"trendyol/internal/domain/item"
	"trendyol/internal/domain/promotion"
)

type MockCartQuery struct {
	items []item.Item
	price float64
}

func (m *MockCartQuery) Items() []item.Item {
	return m.items
}

func (m *MockCartQuery) Price() float64 {
	return m.price
}

func TestFindBestPromotion(t *testing.T) {
	appliers := promotion.NewPromotionAppliers()

	testCases := []struct {
		items            []item.Item
		totalPrice       float64
		expectedID       int
		expectedDiscount float64
	}{
		// category promotion
		{
			items: []item.Item{
				item.NewItem(1, 3003, 1, 1, 25000.0),
				item.NewItem(2, 3003, 2, 1, 25000.0),
			},
			totalPrice:       50000,
			expectedID:       5676,
			expectedDiscount: 2500.0,
		},
		// total price promotion
		{
			items: []item.Item{
				item.NewItem(1, 3003, 1, 1, 1000.0),
				item.NewItem(2, 3004, 2, 1, 40000.0),
			},
			totalPrice:       41000,
			expectedID:       1232,
			expectedDiscount: 1000.0,
		},
		// same seller promotion
		{
			items: []item.Item{
				item.NewItem(1, 3004, 1, 1, 1000.0),
				item.NewItem(2, 3004, 1, 1, 4000.0),
				item.NewItem(3, 3004, 1, 1, 40000.0),
			},
			totalPrice:       45000,
			expectedID:       9909,
			expectedDiscount: 4500.0,
		},
		// mixed promotions
		{
			items: []item.Item{
				item.NewItem(1, 3003, 1, 1, 1000.0),  // Category Promotion applicable
				item.NewItem(2, 3004, 1, 1, 1000.0),  // Same Seller Promotion applicable
				item.NewItem(3, 3004, 1, 1, 40000.0), // Same Seller Promotion applicable
			},
			totalPrice:       42000, // Total Price Promotion applicable
			expectedID:       9909,  // Should choose Same Seller Promotion since it will give the highest discount of 4200.0
			expectedDiscount: 4200.0,
		},
		// mixed promotions
		{
			items: []item.Item{
				item.NewItem(1, 3003, 1, 1, 2500.0), // Category Promotion applicable
				item.NewItem(2, 3004, 2, 1, 2500.0), // Different Seller
				item.NewItem(3, 3005, 3, 1, 4000.0), // Different Seller
			},
			totalPrice:       9000, // Total Price Promotion applicable
			expectedID:       1232, // Should choose Total Price Promotion since it will give the highest discount of 500.0
			expectedDiscount: 500.0,
		},
	}

	for _, tc := range testCases {
		mockCart := &MockCartQuery{
			items: tc.items,
			price: tc.totalPrice,
		}

		promotionService := NewPromotionService(appliers)
		bestPromotion := promotionService.FindBestPromotion(mockCart)

		if bestPromotion.Id() != tc.expectedID {
			t.Errorf("Expected best promotion id to be %d, got %d", tc.expectedID, bestPromotion.Id())
		}

		if bestPromotion.Discount() != tc.expectedDiscount {
			t.Errorf("Expected best promotion discount to be %f, got %f", tc.expectedDiscount, bestPromotion.Discount())
		}
	}
}
