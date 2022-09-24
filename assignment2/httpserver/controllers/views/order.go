package views

import (
	"time"
)

type ItemsOrderGet struct {
	ID          int    `json:"item_id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

type OrderGet struct {
	ID           int             `json:"order_id"`
	CustomerName string          `json:"customer_name"`
	OrderedAt    time.Time       `json:"ordered_at"`
	Items        []ItemsOrderGet `json:"items"`
}

type OrderDelete struct {
	ID           int    `json:"order_id"`
	CustomerName string `json:"customer_name"`
}
