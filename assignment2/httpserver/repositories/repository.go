package repositories

import (
	"github.com/rifqoi/hacktiv8-assignments/assignment2/httpserver/repositories/models"
)

type OrderRepo interface {
	GetOrders(limit int) (*[]models.Order, error)
	CreateOrder(order *models.Order) error
	UpdateOrder(order *models.Order) error
	DeleteOrderByID(id int, model *models.Order) error
}

type ItemRepo interface {
	CreateItem(item *models.Item) error
}
