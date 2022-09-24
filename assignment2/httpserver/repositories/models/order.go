package models

import "time"

type Order struct {
	ID           int `gorm:"primaryKey;autoIncrement"`
	CustomerName string
	OrderedAt    time.Time
	Items        []Item `json:"items" gorm:"foreignKey:OrderID"`
}
