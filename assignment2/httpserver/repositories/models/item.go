package models

type Item struct {
	ID          int `gorm:"primaryKey;autoIncrement"`
	ItemCode    string
	Description string
	Quantity    int
	OrderID     int
	Order       *Order `json:",omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
