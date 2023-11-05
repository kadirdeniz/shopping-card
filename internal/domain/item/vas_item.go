package item

import "errors"

const (
	MaxVasItem        = 3
	VasItemCategoryId = 3242
	VasItemSellerId   = 5003
)

var (
	ErrMaxQuantityForVasItem = errors.New("max quantity for vas item is 3")
	ErrInvalidCategoryId     = errors.New("invalid category id")
	ErrInvalidSellerId       = errors.New("invalid seller id")
)

type vasItem struct {
	item
}

func (i *vasItem) Validate() error {
	if i.CategoryId() != VasItemCategoryId {
		return ErrInvalidCategoryId
	}

	if i.SellerId() != VasItemSellerId {
		return ErrInvalidSellerId
	}

	return nil
}
