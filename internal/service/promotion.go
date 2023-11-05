package service

import (
	"trendyol/internal/domain"
	"trendyol/internal/domain/promotion"
)

type PromotionService interface {
	FindBestPromotion(cart domain.CartQuery) promotion.Promotion
}

type promotionService struct {
	appliers []promotion.PromotionApplier
}

func NewPromotionService(appliers []promotion.PromotionApplier) PromotionService {
	return &promotionService{appliers: appliers}
}

// Multiple promotions are not accepted for the same cart.
// In case of multiple promotions, the best one (providing the maximum discount) is selected for the customer.
func (s *promotionService) FindBestPromotion(cart domain.CartQuery) promotion.Promotion {
	bestPromotion := promotion.NewPromotion(0, 0.0)

	for _, applier := range s.appliers {
		promotion := applier.Apply(cart)
		if promotion.Discount() > bestPromotion.Discount() {
			bestPromotion = promotion
		}
	}

	return bestPromotion
}
