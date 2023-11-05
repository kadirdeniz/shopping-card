package item

import (
	"testing"
)

func TestItemMethods(t *testing.T) {
	testCases := []struct {
		id         int
		categoryId int
		sellerId   int
		quantity   int
		price      float64
		totalPrice float64
		isEmpty    bool
	}{
		{1, DigitalItemCategoryId, 3, 4, 5.0, 20.0, false},
		{6, VasItemCategoryId, 8, 9, 10.0, 90.0, false},
		{11, 12, 13, 14, 15.0, 210.0, false},
		{0, 0, 0, 0, 0.0, 0.0, true},
	}

	for _, tc := range testCases {
		itemInstance := item{id: tc.id, categoryId: tc.categoryId, sellerId: tc.sellerId, price: tc.price, quantity: tc.quantity}

		// Test Id()
		if itemInstance.Id() != tc.id {
			t.Errorf("Expected Id %d, got %d", tc.id, itemInstance.Id())
		}

		// Test CategoryId()
		if itemInstance.CategoryId() != tc.categoryId {
			t.Errorf("Expected CategoryId %d, got %d", tc.categoryId, itemInstance.CategoryId())
		}

		// Test SellerId()
		if itemInstance.SellerId() != tc.sellerId {
			t.Errorf("Expected SellerId %d, got %d", tc.sellerId, itemInstance.SellerId())
		}

		// Test Price()
		if itemInstance.Price() != tc.price {
			t.Errorf("Expected Price %.2f, got %.2f", tc.price, itemInstance.Price())
		}

		// Test Quantity()
		if itemInstance.Quantity() != tc.quantity {
			t.Errorf("Expected Quantity %d, got %d", tc.quantity, itemInstance.Quantity())
		}

		// Test TotalPrice()
		if itemInstance.TotalPrice() != tc.totalPrice {
			t.Errorf("Expected TotalPrice %.2f, got %.2f", tc.totalPrice, itemInstance.TotalPrice())
		}

		// Test IsEmpty()
		if itemInstance.IsEmpty() != tc.isEmpty {
			t.Errorf("Expected IsEmpty to be %v, got %v", tc.isEmpty, itemInstance.IsEmpty())
		}
	}
}

func TestNewItem(t *testing.T) {
	testCases := []struct {
		id           int
		categoryId   int
		sellerId     int
		quantity     int
		price        float64
		totalPrice   float64
		expectedType string
	}{
		{1, DigitalItemCategoryId, 3, 4, 5.0, 20.0, "digitalItem"},
		{6, VasItemCategoryId, 8, 9, 10.0, 90.0, "vasItem"},
		{11, 12, 13, 14, 15.0, 210.0, "defaultItem"},
		{0, 0, 0, 0, 0.0, 0.0, "defaultItem"},
	}

	for _, tc := range testCases {
		item := NewItem(tc.id, tc.categoryId, tc.sellerId, tc.quantity, tc.price)
		var typeName string

		switch item.(type) {
		case *digitalItem:
			typeName = "digitalItem"
		case *vasItem:
			typeName = "vasItem"
		case *defaultItem:
			typeName = "defaultItem"
		case nil:
			typeName = "nil"
		default:
			typeName = "unknown"
		}

		if typeName != tc.expectedType {
			t.Errorf("Type mismatch, expected %s got %s", tc.expectedType, typeName)
		}
	}
}
