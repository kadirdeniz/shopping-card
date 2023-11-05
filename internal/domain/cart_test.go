package domain

import (
	"testing"
	"trendyol/internal/domain/item"
)

func NewTestMockCart() cart {
	return cart{
		items: []item.Item{
			item.NewItem(1, item.FurnitureCategoryId, 2, 1, 100.0),
			item.NewItem(2, item.DigitalItemCategoryId, 1, 2, 200.0),
		},
		price: 500.0,
	}
}

func TestAddItem_MaxQuantityForAddItem(t *testing.T) {
	cart := cart{}

	err := cart.AddItem(item.NewItem(1, 3003, 0, 20, 100.0))
	if err == nil || err != item.ErrMaxQuantityForDefaultItem {
		t.Errorf("Expected ErrMaxQuantityForVasItem, got %s", err)
	}
}

func TestAddItem_MaxUniqueItems(t *testing.T) {
	cart := cart{
		items: []item.Item{
			item.NewItem(1, 3003, 1, 1, 100.0),
			item.NewItem(2, 3004, 1, 1, 200.0),
			item.NewItem(3, 3005, 1, 1, 300.0),
			item.NewItem(4, 3006, 1, 1, 400.0),
			item.NewItem(5, 3007, 1, 1, 500.0),
			item.NewItem(6, 3008, 1, 1, 600.0),
			item.NewItem(7, 3009, 1, 1, 700.0),
			item.NewItem(8, 3010, 1, 1, 800.0),
			item.NewItem(9, 3011, 1, 1, 900.0),
			item.NewItem(10, 3012, 1, 1, 1000.0),
		},
		price: 5500.0,
	}

	// Try adding one more unique item
	err := cart.AddItem(item.NewItem(MaxUniqueItems, 3004, 1, 1, 100.0))
	if err != ErrMaxItem {
		t.Errorf("Expected ErrMaxItem, got %s", err)
	}
}

func TestAddItem_MaxTotalItems(t *testing.T) {
	cart := cart{}
	for i := 0; i < MaxUniqueItems-1; i++ {
		err := cart.AddItem(item.NewItem(1, 3003, 1, 3, 100.0))
		if err != nil {
			t.Errorf("Expected nil, got %s", err)
		}
	}

	// Try adding one more item
	err := cart.AddItem(item.NewItem(1, 3003, 1, 5, 100.0))
	if err != ErrMaxTotalNumberOfItems {
		t.Errorf("Expected ErrMaxTotalNumberOfItems, got %s", err)
	}
}

func TestAddItem_MaxTotalPrice(t *testing.T) {
	cart := cart{}

	err := cart.AddItem(item.NewItem(1, 3003, 1, 1, 500000.0))
	if err != nil {
		t.Errorf("Expected nil, got %s", err)
	}

	err = cart.AddItem(item.NewItem(1, 3003, 1, 1, 100.0))
	if err != ErrMaxTotalPrice {
		t.Errorf("Expected ErrMaxTotalPrice, got %s", err)
	}
}

func TestAddItem_CheckPrice(t *testing.T) {
	cart := cart{}

	err := cart.AddItem(item.NewItem(1, 3003, 1, 1, 100.0))
	if err != nil {
		t.Errorf("Expected nil, got %s", err)
	}

	if cart.price != 100.0 {
		t.Errorf("Expected 100.0, got %f", cart.Price())
	}

	err = cart.AddItem(item.NewItem(2, 3004, 1, 2, 200.0))
	if err != nil {
		t.Errorf("Expected nil, got %s", err)
	}

	if cart.price != 500.0 {
		t.Errorf("Expected 300.0, got %f", cart.Price())
	}
}

func TestAddItem_CheckItems(t *testing.T) {
	cart := cart{}

	err := cart.AddItem(item.NewItem(1, 3003, 1, 1, 100.0))
	if err != nil {
		t.Errorf("Expected nil, got %s", err)
	}

	if len(cart.Items()) != 1 {
		t.Errorf("Expected 1 item, got %d", len(cart.Items()))
	}

	err = cart.AddItem(item.NewItem(2, 3004, 1, 2, 200.0))
	if err != nil {
		t.Errorf("Expected nil, got %s", err)
	}

	if len(cart.Items()) != 2 {
		t.Errorf("Expected 2 items, got %d", len(cart.Items()))
	}
}

func TestAddVasItem_ItemNotFound(t *testing.T) {
	cart := NewTestMockCart()

	err := cart.AddVasItem(item.NewItem(1, 3003, 1, 1, 100.0), 3)
	if err != ErrItemNotFound {
		t.Errorf("Expected ErrItemNotFound, got %s", err)
	}
}

func TestAddVasItem_AddedSuccessfully(t *testing.T) {
	cart := NewTestMockCart()

	err := cart.AddVasItem(item.NewItem(1, item.VasItemCategoryId, item.VasItemSellerId, 1, 10.0), 1)
	if err != nil {
		t.Errorf("Expected nil, got %s", err)
	}
}

func TestAddVasItem_InvalidItem(t *testing.T) {
	cart := NewTestMockCart()

	err := cart.AddVasItem(item.NewItem(1, item.VasItemCategoryId, item.VasItemSellerId, 1, 100.0), 2)
	if err == nil {
		t.Errorf("Expected InvalidItem, got %s", err)
	}
}

func TestAddVasItem_InvalidVasItem(t *testing.T) {
	cart := NewTestMockCart()

	err := cart.AddVasItem(item.NewItem(1, item.VasItemCategoryId, 2, 1, 100.0), 1)
	if err == nil {
		t.Errorf("Expected ErrInvalidVasItem, got %s", err)
	}
}

func TestAddVasItem_InvalidVasItemQuantity(t *testing.T) {
	cart := NewTestMockCart()

	err := cart.AddVasItem(item.NewItem(1, item.VasItemCategoryId, item.VasItemSellerId, 1, 1000.0), 1)
	if err == nil {
		t.Errorf("Expected ErrInvalidVasItem, got %s", err)
	}
}

func TestAddVasItem_CheckPrice(t *testing.T) {
	cart := &cart{
		items: []item.Item{
			item.NewItem(1, item.FurnitureCategoryId, 2, 1, 500.0),
		},
		price: 500.0,
	}

	// Try adding a vas item to an existing item
	err := cart.AddVasItem(item.NewItem(1, item.VasItemCategoryId, item.VasItemSellerId, 1, 10.0), 1)
	if err != nil {
		t.Errorf("Expected nil, got %s", err)
	}

	if cart.Price() != 510.0 {
		t.Errorf("Expected 510.0, got %f", cart.Price())
	}
}

func TestRemoveItem_ItemNotFound(t *testing.T) {
	cart := NewTestMockCart()

	// Remove the item
	err := cart.RemoveItem(3)
	if err != ErrItemNotFound {
		t.Errorf("Expected nil, got %s", err)
	}
}

func TestRemoveItem_CheckPrice(t *testing.T) {
	cart := NewTestMockCart()

	// Remove the item
	err := cart.RemoveItem(1)
	if err != nil {
		t.Errorf("Expected nil, got %s", err)
	}

	// Check the price
	if cart.Price() != 400.0 {
		t.Errorf("Expected 400.0, got %f", cart.Price())
	}
}

func TestRemoveItem_CheckItems(t *testing.T) {
	cart := NewTestMockCart()

	// Remove the item
	err := cart.RemoveItem(1)
	if err != nil {
		t.Errorf("Expected nil, got %s", err)
	}

	// Check the items
	if len(cart.Items()) != 1 {
		t.Errorf("Expected 1 items, got %d", len(cart.Items()))
	}

	// Check the price
	if cart.Price() != 400.0 {
		t.Errorf("Expected 400.0, got %f", cart.Price())
	}
}

func TestReset_ClearsCartState(t *testing.T) {
	cart := NewTestMockCart()

	// Reset the cart
	cart.Reset()

	// Check that the cart is empty
	if len(cart.Items()) != 0 {
		t.Errorf("Expected cart to be empty, got %d items", len(cart.Items()))
	}

	// Check that the price is zero
	if cart.Price() != 0 {
		t.Errorf("Expected cart to have a total price of 0, got %f", cart.Price())
	}
}

func TestItems_ReturnsItemsInCart(t *testing.T) {
	cart := NewTestMockCart()
	// Check if Items() returns the correct number of items
	if len(cart.Items()) != 2 {
		t.Errorf("Expected 2 items, got %d items", len(cart.Items()))
	}

	// Check if Items() returns the correct items
	if cart.Items()[0].Id() != 1 || cart.Items()[1].Id() != 2 {
		t.Errorf("Items returned are not what were added to the cart")
	}
}

func TestPrice_ReturnsTotalPriceOfItemsInCart(t *testing.T) {
	cart := NewTestMockCart()

	if cart.Price() != 500.0 {
		t.Errorf("Expected total price to be %f, got %f", 500.0, cart.Price())
	}
}

func TestNewCart(t *testing.T) {
	cart := NewCart()

	if len(cart.Items()) != 0 {
		t.Errorf("Expected cart to be empty, got %d items", len(cart.Items()))
	}

	if cart.Price() != 0 {
		t.Errorf("Expected cart to have a total price of 0, got %f", cart.Price())
	}
}
