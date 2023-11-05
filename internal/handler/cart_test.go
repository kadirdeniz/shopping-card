package handler

import (
	"testing"
	"trendyol/internal/domain"
	"trendyol/internal/domain/item"
	"trendyol/internal/domain/promotion"
	"trendyol/internal/dto"
	"trendyol/internal/service"
)

var mockDefaultItemPayload = dto.Payload{
	ItemID:     1,
	CategoryID: item.FurnitureCategoryId,
	SellerID:   1,
	Price:      100.0,
	Quantity:   1,
}

func NewTestCartHandler() CartHandler {
	mockCart := domain.NewCart()
	mockPromotionService := service.NewPromotionService(promotion.NewPromotionAppliers())
	return NewCart(mockCart, mockPromotionService)
}

func TestAddItemToCart(t *testing.T) {
	h := NewTestCartHandler()
	testCases := []struct {
		name     string
		payload  dto.Payload
		expected bool
		message  string
	}{
		{"Success", dto.Payload{1, 1, 1, 100.0, 1, 0}, true, ItemAddedSuccessfullyMsg},
		{"ExceedMaxPrice", dto.Payload{1, 1, 1, 1, 31, 0}, false, "total items limit exceeded"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resp := h.AddItemToCart(tc.payload)
			if resp.Result != tc.expected {
				t.Errorf("Expected Result=%v, got %v", tc.expected, resp.Result)
			}
			if resp.Message != tc.message {
				t.Errorf("Expected Message=%s, got %s", tc.message, resp.Message)
			}
		})
	}
}

func TestAddVasItemToItem(t *testing.T) {
	h := NewTestCartHandler()

	const VasItemAddedSuccessfullyMsg = "Vas item added to item successfully"
	const InvalidSellerIdMsg = "invalid seller id"

	defaultItemPayload := mockDefaultItemPayload
	h.AddItemToCart(defaultItemPayload)

	tests := []struct {
		name     string
		payload  dto.Payload
		expected bool
		message  string
	}{
		{
			"Add VasItem Success",
			dto.Payload{
				VasItemId:  2,
				CategoryID: item.VasItemCategoryId,
				SellerID:   item.VasItemSellerId,
				Price:      10.0,
				Quantity:   1,
				ItemID:     1,
			},
			true,
			VasItemAddedSuccessfullyMsg,
		},
		{
			"Add VasItem Invalid SellerId",
			dto.Payload{
				VasItemId:  2,
				CategoryID: item.VasItemCategoryId,
				SellerID:   1, // invalid SellerId
				Price:      10.0,
				Quantity:   1,
				ItemID:     1,
			},
			false,
			InvalidSellerIdMsg,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resp := h.AddVasItemToItem(test.payload)
			if resp.Result != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, resp.Result)
			}
			if resp.Message != test.message {
				t.Errorf("Expected %s, got %s", test.message, resp.Message)
			}
		})
	}
}

func TestRemoveItem(t *testing.T) {
	h := NewTestCartHandler()

	const ItemRemovedSuccessfullyMsg = "Item removed from cart successfully"
	const ItemNotFoundMsg = "item not found"

	defaultItemPayload := mockDefaultItemPayload
	h.AddItemToCart(defaultItemPayload)

	tests := []struct {
		name     string
		payload  dto.Payload
		setup    bool
		expected bool
		message  string
	}{
		{
			"Remove Item Success",
			defaultItemPayload,
			true,
			true,
			ItemRemovedSuccessfullyMsg,
		},
		{
			"Remove Item Not Found",
			dto.Payload{
				ItemID:     2,
				CategoryID: item.FurnitureCategoryId,
				SellerID:   1,
				Price:      100.0,
				Quantity:   1,
			},
			false,
			false,
			ItemNotFoundMsg,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.setup {
				h.AddItemToCart(test.payload)
			}

			resp := h.RemoveItem(test.payload)

			if resp.Result != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, resp.Result)
			}
			if resp.Message != test.message {
				t.Errorf("Expected %s, got %s", test.message, resp.Message)
			}
		})
	}
}

func TestReset_Success(t *testing.T) {
	h := NewTestCartHandler()

	defaultItemPayload := mockDefaultItemPayload
	h.AddItemToCart(defaultItemPayload)

	resp := h.Reset()
	if !resp.Result {
		t.Errorf("Expected Result=true, got false")
	}
}

func TestDisplay(t *testing.T) {
	h := NewTestCartHandler()

	defaultItemPayload := mockDefaultItemPayload
	h.AddItemToCart(defaultItemPayload)

	h.AddVasItemToItem(dto.Payload{
		ItemID:     1,
		VasItemId:  1,
		CategoryID: item.VasItemCategoryId,
		SellerID:   item.VasItemSellerId,
		Price:      10.0,
		Quantity:   1,
	})

	resp := h.Display()
	if !resp.Result {
		t.Errorf("Expected Result=true, got false")
	}

	cart := resp.Message.(dto.CartResponse)
	if len(cart.Items) != 1 {
		t.Errorf("Expected 1 item, got %d", len(cart.Items))
	}

	if len(cart.Items[0].ItemVases) != 1 {
		t.Errorf("Expected 1 vas item, got %d", len(cart.Items[0].ItemVases))
	}

	if cart.Items[0].ItemVases[0].VasItemId != 1 {
		t.Errorf("Expected vas item id=1, got %d", cart.Items[0].ItemVases[0].VasItemId)
	}

	if cart.Items[0].ItemVases[0].ItemID != 0 {
		t.Errorf("Expected vas item id=1, got %d", cart.Items[0].ItemVases[0].ItemID)
	}

	if cart.TotalPrice != 110.0 {
		t.Errorf("Expected price=110.0, got %f", cart.TotalPrice)
	}

	if cart.TotalDiscount != 250.0 {
		t.Errorf("Expected discount=250.0, got %f", cart.TotalDiscount)
	}

	if cart.AppliedPromotionId != promotion.TotalPricePromotionId {
		t.Errorf("Expected promotionId=1, got %d", cart.AppliedPromotionId)
	}
}

func TestDisplay_EmptyCart(t *testing.T) {
	h := NewTestCartHandler()

	resp := h.Display()
	if !resp.Result {
		t.Errorf("Expected Result=true, got false")
	}

	cart := resp.Message.(dto.CartResponse)
	if len(cart.Items) != 0 {
		t.Errorf("Expected 0 item, got %d", len(cart.Items))
	}

	if cart.TotalPrice != 0.0 {
		t.Errorf("Expected price=0.0, got %f", cart.TotalPrice)
	}

	if cart.TotalDiscount != 0.0 {
		t.Errorf("Expected discount=0.0, got %f", cart.TotalDiscount)
	}

	if cart.AppliedPromotionId != 0 {
		t.Errorf("Expected promotionId=0, got %d", cart.AppliedPromotionId)
	}
}
