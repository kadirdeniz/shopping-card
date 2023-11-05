package promotion

import (
	"testing"
	"trendyol/internal/domain/item"
)

type MockCart struct {
	items []item.Item
	price float64
}

func (m *MockCart) Items() []item.Item {
	return m.items
}

func (m *MockCart) Price() float64 {
	return m.price
}

func TestPromotionMethods(t *testing.T) {
	testCases := []struct {
		id       int
		discount float64
	}{
		{1, 10.5},
		{2, 20.0},
		{3, 0.0},
	}

	for _, tc := range testCases {
		p := NewPromotion(tc.id, tc.discount)

		if p.Id() != tc.id {
			t.Errorf("ID mismatch, expected %d got %d", tc.id, p.Id())
		}

		if p.Discount() != tc.discount {
			t.Errorf("Discount mismatch, expected %f got %f", tc.discount, p.Discount())
		}
	}
}

func TestNewPromotionAppliers(t *testing.T) {
	promotionAppliers := NewPromotionAppliers()

	if len(promotionAppliers) != 3 {
		t.Errorf("Expected 3 promotion appliers, got %d", len(promotionAppliers))
	}
}
