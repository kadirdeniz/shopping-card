package promotion

import (
	"testing"
)

func TestTotalPricePromotionApply(t *testing.T) {
	testCases := []struct {
		totalPrice   float64
		expectedID   int
		expectedDisc float64
	}{
		{
			totalPrice:   4000,
			expectedID:   TotalPricePromotionId,
			expectedDisc: DiscountForEconomicRange,
		},
		{
			totalPrice:   7000,
			expectedID:   TotalPricePromotionId,
			expectedDisc: DiscountForMidRange,
		},
		{
			totalPrice:   30000,
			expectedID:   TotalPricePromotionId,
			expectedDisc: DiscountForPremiumRange,
		},
		{
			totalPrice:   60000,
			expectedID:   TotalPricePromotionId,
			expectedDisc: DiscountForLuxuryRange,
		},
	}

	for _, tc := range testCases {
		mockCart := &MockCart{
			price: tc.totalPrice,
		}

		prom := NewTotalPricePromotion()
		appliedProm := prom.Apply(mockCart)

		if appliedProm.Id() != tc.expectedID {
			t.Errorf("Expected promotion id to be %d, got %d", tc.expectedID, appliedProm.Id())
		}

		if appliedProm.Discount() != tc.expectedDisc {
			t.Errorf("Expected discount to be %f, got %f", tc.expectedDisc, appliedProm.Discount())
		}
	}
}
