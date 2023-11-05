package handler

import (
	"trendyol/internal/domain"
	"trendyol/internal/domain/item"
	"trendyol/internal/dto"
	"trendyol/internal/service"
)

const (
	ItemAddedSuccessfullyMsg     = "Item added to cart successfully"
	VasItemAddedSuccessfullyMsg  = "Vas item added to item successfully"
	ItemRemovedSuccessfullyMsg   = "Item removed from cart successfully"
	CartResetSuccessfullyMsg     = "Cart reset successfully"
	CartDisplayedSuccessfullyMsg = "Cart displayed successfully"
)

type CartHandler interface {
	AddItemToCart(request dto.Payload) dto.Response
	AddVasItemToItem(request dto.Payload) dto.Response
	RemoveItem(request dto.Payload) dto.Response
	Reset() dto.Response
	Display() dto.Response
}

type cartHandler struct {
	cart             domain.Cart
	promotionService service.PromotionService
}

func NewCart(cart domain.Cart, promotionService service.PromotionService) CartHandler {
	return &cartHandler{cart: cart, promotionService: promotionService}
}

func (h *cartHandler) AddItemToCart(request dto.Payload) dto.Response {
	item := item.NewItem(request.ItemID, request.CategoryID, request.SellerID, request.Quantity, request.Price)
	if err := h.cart.AddItem(item); err != nil {
		return dto.Response{
			Result:  false,
			Message: err.Error(),
		}
	}

	return dto.Response{
		Result:  true,
		Message: ItemAddedSuccessfullyMsg,
	}
}

func (h *cartHandler) AddVasItemToItem(request dto.Payload) dto.Response {
	vasItem := item.NewItem(request.VasItemId, request.CategoryID, request.SellerID, request.Quantity, request.Price)
	if err := h.cart.AddVasItem(vasItem, request.ItemID); err != nil {
		return dto.Response{
			Result:  false,
			Message: err.Error(),
		}
	}

	return dto.Response{
		Result:  true,
		Message: VasItemAddedSuccessfullyMsg,
	}
}

func (h *cartHandler) RemoveItem(request dto.Payload) dto.Response {
	if err := h.cart.RemoveItem(request.ItemID); err != nil {
		return dto.Response{
			Result:  false,
			Message: err.Error(),
		}
	}

	return dto.Response{
		Result:  true,
		Message: ItemRemovedSuccessfullyMsg,
	}
}

func (h *cartHandler) Reset() dto.Response {
	h.cart.Reset()

	return dto.Response{
		Result:  true,
		Message: CartResetSuccessfullyMsg,
	}
}

func (h *cartHandler) Display() dto.Response {
	cart := h.cart
	bestPromotion := h.promotionService.FindBestPromotion(cart)

	response := dto.CartResponse{
		TotalPrice:         cart.Price(),
		AppliedPromotionId: bestPromotion.Id(),
		TotalDiscount:      bestPromotion.Discount(),
	}

	for _, itemObj := range cart.Items() {
		itemPayload := toItemPayload(itemObj)

		if defaultItem, ok := itemObj.(item.DefaultItem); ok {
			for _, vasItemObj := range defaultItem.VasItems() {
				vasItemPayload := toPayload(vasItemObj)
				itemPayload.ItemVases = append(itemPayload.ItemVases, vasItemPayload)
			}
		}

		response.Items = append(response.Items, itemPayload)
	}

	return dto.Response{
		Result:  true,
		Message: response,
	}
}

func toItemPayload(itemObj item.Item) dto.ItemPayload {
	return dto.ItemPayload{
		ItemID:     itemObj.Id(),
		CategoryID: itemObj.CategoryId(),
		SellerID:   itemObj.SellerId(),
		Price:      itemObj.Price(),
		Quantity:   itemObj.Quantity(),
	}
}

func toPayload(vasItemObj item.Item) dto.Payload {
	return dto.Payload{
		VasItemId:  vasItemObj.Id(),
		CategoryID: vasItemObj.CategoryId(),
		SellerID:   vasItemObj.SellerId(),
		Price:      vasItemObj.Price(),
		Quantity:   vasItemObj.Quantity(),
	}
}
