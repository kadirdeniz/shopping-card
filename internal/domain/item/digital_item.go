package item

import "errors"

const (
	MaxQuantityForDigitalItem = 5
	DigitalItemCategoryId     = 7889
)

var (
	ErrMaxQuantityForDigitalItem = errors.New("max quantity for digital item is 5")
)

type digitalItem struct {
	item
}

func (i *digitalItem) Validate() error {
	if i.Quantity() >= MaxQuantityForDigitalItem {
		return ErrMaxQuantityForDigitalItem
	}

	return nil
}
