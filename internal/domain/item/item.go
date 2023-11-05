package item

type Item interface {
	Id() int
	CategoryId() int
	SellerId() int
	Price() float64
	Quantity() int
	Validate() error
	TotalPrice() float64
	IsEmpty() bool
}

type item struct {
	id         int
	categoryId int
	sellerId   int
	price      float64
	quantity   int
}

func NewItem(id, categoryId, sellerId, quantity int, price float64) Item {
	baseItem := item{
		id:         id,
		categoryId: categoryId,
		sellerId:   sellerId,
		price:      price,
		quantity:   quantity,
	}

	switch categoryId {
	case DigitalItemCategoryId:
		return &digitalItem{
			item: baseItem,
		}
	case VasItemCategoryId:
		return &vasItem{
			item: baseItem,
		}
	default:
		return &defaultItem{
			item:     baseItem,
			vasItems: []vasItem{},
		}
	}
}

func (i *item) Id() int {
	return i.id
}

func (i *item) CategoryId() int {
	return i.categoryId
}

func (i *item) SellerId() int {
	return i.sellerId
}

func (i *item) Price() float64 {
	return i.price
}

func (i *item) Quantity() int {
	return i.quantity
}

func (i *item) IsEmpty() bool {
	return i.id == 0 || i == nil
}

func (i *item) TotalPrice() float64 {
	return float64(i.Quantity()) * i.Price()
}
