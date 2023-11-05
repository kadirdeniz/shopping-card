package dto

type ItemPayload struct {
	ItemID     int     `json:"itemId"`
	CategoryID int     `json:"categoryId"`
	SellerID   int     `json:"sellerId"`
	Price      float64 `json:"price"`
	Quantity   int     `json:"quantity"`
	ItemVases  []Payload
}

type Payload struct {
	ItemID     int     `json:"itemId"`
	CategoryID int     `json:"categoryId"`
	SellerID   int     `json:"sellerId"`
	Price      float64 `json:"price"`
	Quantity   int     `json:"quantity"`
	VasItemId  int     `json:"vasItemId"`
}

type CartResponse struct {
	TotalPrice         float64       `json:"totalPrice"`
	Items              []ItemPayload `json:"items"`
	AppliedPromotionId int           `json:"appliedPromotionId"`
	TotalDiscount      float64       `json:"totalDiscount"`
}

type Response struct {
	Result  bool        `json:"result"`
	Message interface{} `json:"message"`
}

type Request struct {
	Command string  `json:"command"`
	Payload Payload `json:"payload"`
}
