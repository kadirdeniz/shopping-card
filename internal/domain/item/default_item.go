package item

import "errors"

const (
	MaxQuantityForDefaultItem = 10
	FurnitureCategoryId       = 1001
	ElectronicsCategoryId     = 3004
)

var (
	ErrMaxQuantityForDefaultItem    = errors.New("max quantity for default item is 10")
	ErrInvalidVasItemPrice          = errors.New("invalid vas item price")
	ErrInvalidVasItemType           = errors.New("invalid vas item type")
	ErrInvalidDefaultItemCategoryId = errors.New("invalid default item category id")
)

type DefaultItem interface {
	Item
	VasItems() []Item
	AddVasItem(vasItem Item) error
}

type defaultItem struct {
	item
	vasItems []vasItem
}

func NewDefaultItem(itemObj Item) DefaultItem {
	return &defaultItem{
		item: item{
			id:         itemObj.Id(),
			categoryId: itemObj.CategoryId(),
			sellerId:   itemObj.SellerId(),
			price:      itemObj.Price(),
			quantity:   itemObj.Quantity(),
		},
		vasItems: []vasItem{},
	}
}

func (i *defaultItem) AddVasItem(vasItemObj Item) error {
	if err := vasItemObj.Validate(); err != nil {
		return err
	}

	if vasItemObj.TotalPrice() > i.Price()*float64(i.Quantity()) {
		return ErrInvalidVasItemPrice
	}

	if len(i.vasItems) >= MaxVasItem {
		return ErrMaxQuantityForVasItem
	}

	if i.categoryId != FurnitureCategoryId && i.categoryId != ElectronicsCategoryId {
		return ErrInvalidDefaultItemCategoryId
	}

	// only vasItem can be added to default item
	vasItem, ok := vasItemObj.(*vasItem)
	if !ok {
		return ErrInvalidVasItemType
	}

	i.vasItems = append(i.vasItems, *vasItem)

	return nil
}

func (i *defaultItem) Validate() error {
	if i.Quantity() >= MaxQuantityForDefaultItem {
		return ErrMaxQuantityForDefaultItem
	}
	return nil
}

// default item price + vas item price
func (i *defaultItem) TotalPrice() float64 {
	totalPrice := float64(i.Quantity()) * i.Price()

	for _, vas := range i.vasItems {
		totalPrice += vas.TotalPrice()
	}

	return totalPrice
}

// return []vasItem as []Item
func (i *defaultItem) VasItems() []Item {
	var items []Item

	for _, vas := range i.vasItems {
		items = append(items, &vas)
	}

	return items
}
