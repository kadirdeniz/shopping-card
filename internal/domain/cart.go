package domain

import (
	"errors"
	"trendyol/internal/domain/item"
)

const (
	MaxUniqueItems = 10
	MaxTotalItems  = 30
	MaxTotalPrice  = 500000.0 // TL
)

var (
	ErrMaxItem               = errors.New("maximum unique items limit exceeded")
	ErrMaxTotalPrice         = errors.New("total price limit exceeded")
	ErrMaxTotalNumberOfItems = errors.New("total items limit exceeded")
	ErrValidationItem        = errors.New("item validation failed")
	ErrItemNotFound          = errors.New("item not found")
	ErrInvalidDefaultItem    = errors.New("invalid default item")
)

type Cart interface {
	AddItem(item item.Item) error
	AddVasItem(vasItem item.Item, itemId int) error
	RemoveItem(itemId int) error
	Reset()
	CartQuery
}

type CartQuery interface {
	Items() []item.Item
	Price() float64
}

type cart struct {
	items []item.Item
	price float64
}

func NewCart() Cart {
	return &cart{}
}

func (c *cart) AddItem(newItem item.Item) error {
	// Validate the item and the current state of the cart
	if err := c.validate(newItem); err != nil {
		return err
	}

	if len(c.items) >= MaxUniqueItems {
		return ErrMaxItem
	}

	// Add the item to the cart
	c.items = append(c.items, newItem)

	// Update the total price of the cart
	c.price += newItem.TotalPrice()

	return nil
}

func (c *cart) AddVasItem(vasItem item.Item, itemId int) error {
	if err := c.validate(vasItem); err != nil {
		return err
	}

	// VasItem can be added to the cart only if 
	// the cart has the item with the given itemId
	itemObj := c.getItemById(itemId)
	if itemObj.IsEmpty() {
		return ErrItemNotFound
	}

	// Add the vas item to the default item
	// vasItem only can be added to default item
	// with CategoryId FurnitureCategoryId or ElectronicsCategoryId
	defaultItem, ok := itemObj.(item.DefaultItem)
	if !ok {
		return ErrInvalidDefaultItem
	}

	err := defaultItem.AddVasItem(vasItem)
	if err != nil {
		return err
	}

	c.price += vasItem.TotalPrice()

	return nil
}

func (c *cart) RemoveItem(itemId int) error {

	// Find the item and remove it from the cart
	for i, cartItem := range c.items {
		if cartItem.Id() == itemId {
			c.price -= cartItem.TotalPrice()
			c.items = append(c.items[:i], c.items[i+1:]...)
			return nil
		}
	}

	return ErrItemNotFound
}

func (c *cart) Reset() {
	c.items = nil
	c.price = 0
}

func (c *cart) Items() []item.Item {
	return c.items
}

func (c *cart) Price() float64 {
	return c.price
}

// validate methods validates the cart and
// the given item
func (c *cart) validate(item item.Item) error {
	if c.price+item.TotalPrice() > MaxTotalPrice {
		return ErrMaxTotalPrice
	}

	if c.totalNumberOfItems()+item.Quantity() > MaxTotalItems {
		return ErrMaxTotalNumberOfItems
	}

	if err := item.Validate(); err != nil {
		return err
	}

	return nil
}

// totalNumberOfItems returns the total number of items in the cart
// vasItems are not included
func (c *cart) totalNumberOfItems() int {
	var total int

	for _, item := range c.items {
		total += item.Quantity()
	}

	return total
}

// getItemById returns the item with the given id
// if the item is not found, it returns an empty item
func (c *cart) getItemById(itemId int) item.Item {
	for _, item := range c.items {
		if item.Id() == itemId {
			return item
		}
	}

	return item.NewItem(0, 0, 0, 0, 0)
}
