package promotion

import (
	"trendyol/internal/domain"
	"trendyol/internal/domain/item"
)

const (
	SameSellerPromotionId = 9909
)

type sameSellerPromotion struct {
	promotion
}

func NewSameSellerPromotion() PromotionApplier {
	return &sameSellerPromotion{}
}

// Apply applies the same seller promotion to the cart
// and returns the promotion.
// It calculates the total discount amount by
// multiplying the price of the cart by 0.1.
func (p *sameSellerPromotion) Apply(cart domain.CartQuery) Promotion {
	sellerIDs := make(map[int]bool)
	totalUniqueItems := len(cart.Items())
	var discount float64
	var id int = SameSellerPromotionId

	for _, itemObj := range cart.Items() {
		if itemObj.CategoryId() != item.VasItemCategoryId {
			sellerIDs[itemObj.SellerId()] = true
		}
	}
	// if the number of unique sellers is greater than the number of items
	// then there are at least one seller with more than one item.
	// in this case, apply the promotion.
	if totalUniqueItems > len(sellerIDs) {
		discount = float64(cart.Price()) * 0.1
	}

	return sameSellerPromotion{
		promotion: promotion{
			id:       id,
			discount: discount,
		},
	}
}
