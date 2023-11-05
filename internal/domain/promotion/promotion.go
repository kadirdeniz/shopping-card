package promotion

import "trendyol/internal/domain"

type Promotion interface {
	Id() int
	Discount() float64
}

type PromotionApplier interface {
	Promotion
	Apply(cart domain.CartQuery) Promotion
}

type promotion struct {
	id       int
	discount float64
}

func NewPromotion(id int, discount float64) Promotion {
	return &promotion{id: id, discount: discount}
}

func NewPromotionAppliers() []PromotionApplier {
	return []PromotionApplier{
		NewCategoryPromotion(),
		NewSameSellerPromotion(),
		NewTotalPricePromotion(),
	}
}

func (p promotion) Id() int {
	return p.id
}

func (p promotion) Discount() float64 {
	return p.discount
}
