package item

import (
	"testing"
)

func TestNewDefaultItem(t *testing.T) {
	t.Run("Create Default Item with using NewItem()", func(t *testing.T) {
		item := NewItem(1, 1, 1, 1, 1.0)
		defaultItem := NewDefaultItem(item)
		if defaultItem == nil {
			t.Errorf("Expected defaultItem, got nil")
		}
	})
}

func TestAddVasItem(t *testing.T) {
	t.Run("Adding VasItem with Incorrect SellerID to Furniture Default Item", func(t *testing.T) {
		defaultItem := defaultItem{item{id: 1, categoryId: FurnitureCategoryId, sellerId: 1, quantity: 1, price: 50.0}, []vasItem{}}
		invalidVasItem := NewItem(1, VasItemCategoryId, 1, 2, 1.0)
		err := defaultItem.AddVasItem(invalidVasItem)
		if err == nil {
			t.Errorf("Expected nil for invalid vasItem, got %v", err)
		}
	})

	t.Run("Adding VasItem with Correct SellerID to Furniture Default Item", func(t *testing.T) {
		defaultItem := defaultItem{item{id: 1, categoryId: FurnitureCategoryId, sellerId: 1, quantity: 1, price: 50.0}, []vasItem{}}
		invalidVasItem := NewItem(1, VasItemCategoryId, VasItemSellerId, 2, 1.0)
		err := defaultItem.AddVasItem(invalidVasItem)
		if err != nil {
			t.Errorf("Expected nil for invalid vasItem, got %v", err)
		}
	})

	t.Run("Adding VasItem with Correct SellerID to Electronics Default Item", func(t *testing.T) {
		defaultItem := defaultItem{item{id: 1, categoryId: ElectronicsCategoryId, sellerId: 1, quantity: 1, price: 50.0}, []vasItem{}}
		invalidVasItem := NewItem(1, VasItemCategoryId, VasItemSellerId, 2, 1.0)
		err := defaultItem.AddVasItem(invalidVasItem)
		if err != nil {
			t.Errorf("Expected nil for invalid vasItem, got %v", err)
		}
	})

	t.Run("Adding VasItem with Incorrect SellerID to Electronics Default Item", func(t *testing.T) {
		defaultItem := defaultItem{item{id: 1, categoryId: ElectronicsCategoryId, sellerId: 1, quantity: 1, price: 50.0}, []vasItem{}}
		invalidVasItem := NewItem(1, VasItemCategoryId, 1, 2, 1.0)
		err := defaultItem.AddVasItem(invalidVasItem)
		if err == nil {
			t.Errorf("Expected nil for invalid vasItem, got %v", err)
		}
	})

	t.Run("Adding VasItem with Correct SellerID to Default Item with Incorrect CategoryID", func(t *testing.T) {
		defaultItem := defaultItem{item{id: 1, categoryId: 1, sellerId: 1, quantity: 1, price: 50.0}, []vasItem{}}
		invalidVasItem := NewItem(1, VasItemCategoryId, VasItemSellerId, 2, 1.0)
		err := defaultItem.AddVasItem(invalidVasItem)
		if err == nil || err != ErrInvalidDefaultItemCategoryId {
			t.Errorf("Expected nil for invalid vasItem, got %v", err)
		}
	})

	t.Run("Adding VasItem with Price Greater Than Default Item", func(t *testing.T) {
		defaultItem := defaultItem{item{id: 1, categoryId: FurnitureCategoryId, sellerId: 1, quantity: 1, price: 50.0}, []vasItem{}}
		expensiveVasItem := NewItem(1, VasItemCategoryId, VasItemSellerId, 2, 100.0)
		err := defaultItem.AddVasItem(expensiveVasItem)
		if err == nil || err != ErrInvalidVasItemPrice {
			t.Errorf("Expected error for expensive vasItem, got %v", err)
		}
	})

	t.Run("Adding VasItem to Default Item", func(t *testing.T) {
		defaultItem := defaultItem{item{id: 1, categoryId: FurnitureCategoryId, sellerId: 1, quantity: 1, price: 50.0}, []vasItem{}}
		vasItem := NewItem(1, VasItemCategoryId, VasItemSellerId, 1, 1.0)
		err := defaultItem.AddVasItem(vasItem)
		if err != nil {
			t.Errorf("Expected nil for valid vasItem, got %v", err)
		}

		if len(defaultItem.vasItems) != 1 {
			t.Errorf("Expected 1 vasItem, got %d", len(defaultItem.vasItems))
		}
	})

	t.Run("Exceeding Maximum VasItem Limit", func(t *testing.T) {
		defaultItem := defaultItem{item{id: 1, categoryId: FurnitureCategoryId, sellerId: 1, quantity: 1, price: 50.0}, []vasItem{}}
		for i := 0; i < MaxVasItem; i++ {
			vasItem := NewItem(i+1, VasItemCategoryId, VasItemSellerId, 1, 1.0)
			_ = defaultItem.AddVasItem(vasItem)
		}
		extraVasItem := NewItem(1, VasItemCategoryId, VasItemSellerId, 1, 1.0)
		err := defaultItem.AddVasItem(extraVasItem)
		if err == nil || err != ErrMaxQuantityForVasItem {
			t.Errorf("Expected error for extra vasItem, got %v", err)
		}
	})

	t.Run("Adding Digital Item to Default Item", func(t *testing.T) {
		defaultItem := defaultItem{item{id: 1, categoryId: FurnitureCategoryId, sellerId: 1, quantity: 1, price: 50.0}, []vasItem{}}
		digitalItem := NewItem(1, DigitalItemCategoryId, 1, 2, 1.0)
		err := defaultItem.AddVasItem(digitalItem)
		if err == nil || err != ErrInvalidVasItemType {
			t.Errorf("Expected error for invalid category or seller, got %v", err)
		}
	})
}

func TestValidate(t *testing.T) {
	t.Run("Test validate defaultItem", func(t *testing.T) {
		defaultItem := defaultItem{item{id: 1, categoryId: FurnitureCategoryId, sellerId: 1, quantity: 1, price: 50.0}, []vasItem{}}
		err := defaultItem.Validate()
		if err != nil {
			t.Errorf("Expected nil for invalid defaultItem, got %v", err)
		}
	})

	t.Run("Test validate defaultItem with invalid quantity", func(t *testing.T) {
		defaultItem := defaultItem{item{id: 1, categoryId: FurnitureCategoryId, sellerId: 1, quantity: MaxQuantityForDefaultItem + 1, price: 50.0}, []vasItem{}}
		err := defaultItem.Validate()
		if err == nil || err != ErrMaxQuantityForDefaultItem {
			t.Errorf("Expected error for invalid defaultItem, got %v", err)
		}
	})
}

func TestTotalPrice(t *testing.T) {
	t.Run("Test total price", func(t *testing.T) {
		defaultItem := defaultItem{item{id: 1, categoryId: FurnitureCategoryId, sellerId: 1, quantity: 1, price: 50.0}, []vasItem{}}
		vasItem := NewItem(1, VasItemCategoryId, VasItemSellerId, 1, 1.0)
		_ = defaultItem.AddVasItem(vasItem)

		expectedTotalPrice := 51.0
		actualTotalPrice := defaultItem.TotalPrice()
		if actualTotalPrice != expectedTotalPrice {
			t.Errorf("Total price mismatch. Expected: %f, Got: %f", expectedTotalPrice, actualTotalPrice)
		}
	})
}

func TestVasItems(t *testing.T) {
	t.Run("Test vasItems", func(t *testing.T) {
		defaultItem := defaultItem{item{id: 1, categoryId: FurnitureCategoryId, sellerId: 1, quantity: 1, price: 50.0}, []vasItem{}}
		vasItem := NewItem(1, VasItemCategoryId, VasItemSellerId, 1, 1.0)
		_ = defaultItem.AddVasItem(vasItem)

		expectedVasItems := []Item{vasItem}
		actualVasItems := defaultItem.VasItems()
		if len(actualVasItems) != len(expectedVasItems) {
			t.Errorf("Expected vasItems length %d, got %d", len(expectedVasItems), len(actualVasItems))
		}
	})
}
