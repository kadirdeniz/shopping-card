package promotion

import "trendyol/internal/domain"

type totalPricePromotion struct {
	promotion
}

const (
	TotalPricePromotionId = 1232
)

const (
	DiscountForEconomicRange = 250
	DiscountForMidRange      = 500
	DiscountForPremiumRange  = 1000
	DiscountForLuxuryRange   = 2000
)

const (
	PriceThresholdEconomic = 5000
	PriceThresholdMid      = 10000
	PriceThresholdPremium  = 50000
)

func NewTotalPricePromotion() PromotionApplier {
	return &totalPricePromotion{}
}

// Apply applies the total price promotion to the cart
// and returns the promotion.
func (p *totalPricePromotion) Apply(cart domain.CartQuery) Promotion {
	var id int = TotalPricePromotionId
	var discount float64

	if cart.Price() > 0 && cart.Price() < PriceThresholdEconomic {
		discount = DiscountForEconomicRange
	} else if cart.Price() >= PriceThresholdEconomic && cart.Price() < PriceThresholdMid {
		discount = DiscountForMidRange
	} else if cart.Price() >= PriceThresholdMid && cart.Price() < PriceThresholdPremium {
		discount = DiscountForPremiumRange
	} else if cart.Price() >= PriceThresholdPremium {
		discount = DiscountForLuxuryRange
	}

	return totalPricePromotion{
		promotion: promotion{
			id:       id,
			discount: discount,
		},
	}
}
