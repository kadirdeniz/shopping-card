package pkg

import (
	"testing"
	"trendyol/internal/dto"
)

type MockCartHandler struct {
	AddItemToCartFn    func(dto.Payload) dto.Response
	AddVasItemToItemFn func(dto.Payload) dto.Response
	RemoveItemFn       func(dto.Payload) dto.Response
	ResetFn            func() dto.Response
	DisplayFn          func() dto.Response
	UnknownFn          func() dto.Response
}

func (m *MockCartHandler) AddItemToCart(request dto.Payload) dto.Response {
	return m.AddItemToCartFn(request)
}

func (m *MockCartHandler) AddVasItemToItem(request dto.Payload) dto.Response {
	return m.AddVasItemToItemFn(request)
}

func (m *MockCartHandler) RemoveItem(request dto.Payload) dto.Response {
	return m.RemoveItemFn(request)
}

func (m *MockCartHandler) Reset() dto.Response {
	return m.ResetFn()
}

func (m *MockCartHandler) Display() dto.Response {
	return m.DisplayFn()
}

func TestRun(t *testing.T) {
	mockHandler := &MockCartHandler{
		RemoveItemFn: func(request dto.Payload) dto.Response {
			if request.ItemID == 0000 {
				return dto.Response{Result: true, Message: "Item removed successfully"}
			}
			return dto.Response{Result: false, Message: "Item not found"}
		},
		AddItemToCartFn: func(request dto.Payload) dto.Response {
			if request.ItemID == 0000 {
				return dto.Response{Result: true, Message: "Item added successfully"}
			}
			return dto.Response{Result: false, Message: "something went wrong"}
		},
		AddVasItemToItemFn: func(request dto.Payload) dto.Response {
			if request.ItemID == 0000 {
				return dto.Response{Result: true, Message: "Vas item added successfully"}
			}
			return dto.Response{Result: false, Message: "something went wrong"}
		},
		ResetFn: func() dto.Response {
			return dto.Response{Result: true, Message: "Cart reset successfully"}
		},
		DisplayFn: func() dto.Response {
			return dto.Response{Result: true, Message: "Cart displayed successfully"}
		},
		UnknownFn: func() dto.Response {
			return dto.Response{Result: false, Message: "Unknown command"}
		},
	}

	executor := NewCommandExecutor(mockHandler)

	requests := []dto.Request{
		{
			Command: CommandRemoveItem,
			Payload: dto.Payload{ItemID: 0000},
		},
		{
			Command: CommandRemoveItem,
			Payload: dto.Payload{ItemID: 0001},
		},
		{
			Command: CommandAddItem,
			Payload: dto.Payload{ItemID: 0000},
		},
		{
			Command: CommandAddItem,
			Payload: dto.Payload{ItemID: 0001},
		},
		{
			Command: CommandAddVasItemToItem,
			Payload: dto.Payload{ItemID: 0000},
		},
		{
			Command: CommandAddVasItemToItem,
			Payload: dto.Payload{ItemID: 0001},
		},
		{
			Command: CommandResetCart,
		},
		{
			Command: CommandDisplayCart,
		},
		{
			Command: CommandUnknown,
		},
	}

	responses := executor.Run(requests)

	if !responses[0].Result || responses[0].Message != "Item removed successfully" {
		t.Errorf("Expected Result=true, Message='Item removed successfully', got Result=%v, Message='%s'", responses[0].Result, responses[0].Message)
	}

	if responses[1].Result || responses[1].Message != "Item not found" {
		t.Errorf("Expected Result=false, Message='Item not found', got Result=%v, Message='%s'", responses[1].Result, responses[1].Message)
	}
}
