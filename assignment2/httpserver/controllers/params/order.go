package params

type itemsCreate struct {
	ItemCode    string `json:"item_code" validate:"required"`
	Description string `json:"description" validate:"required"`
	Quantity    int    `json:"quantity" validate:"required"`
}

type OrderCreateRequest struct {
	OrderedAt    string        `json:"ordered_at" validate:"required"`
	CustomerName string        `json:"customer_name" validate:"required"`
	Items        []itemsCreate `json:"items" validate:"gt=0,required,dive,required"`
}

type itemsUpdate struct {
	LineItemID  int    `json:"line_item_id" validate:"required"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

type OrderUpdateRequest struct {
	OrderID      int           `json:"order_id" validate:"required"`
	CustomerName string        `json:"customer_name" `
	OrderedAt    string        `json:"ordered_at" `
	Items        []itemsUpdate `json:"items" `
}
