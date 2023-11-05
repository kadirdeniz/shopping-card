package promotion

import "trendyol/internal/domain"

const (
	CategoryPromotionId  = 5676
	DiscountedCategoryId = 3003
)

type categoryPromotion struct {
	promotion
}

func NewCategoryPromotion() PromotionApplier {
	return &categoryPromotion{}
}

// Apply applies the category promotion to the cart
// and returns the promotion.
// It calculates the total discount amount by
// multiplying the price of the discounted category
// by 0.05.
func (p *categoryPromotion) Apply(cart domain.CartQuery) Promotion {
	var totalDiscount float64
	var id int = CategoryPromotionId

	for _, itemObj := range cart.Items() {
		if itemObj.CategoryId() == DiscountedCategoryId {
			totalDiscount += (itemObj.Price() * float64(itemObj.Quantity())) * 0.05
		}
	}

	return categoryPromotion{
		promotion: promotion{
			id:       id,
			discount: totalDiscount,
		},
	}
}
