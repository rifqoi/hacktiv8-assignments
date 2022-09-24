package gorm

import (
	"github.com/rifqoi/hacktiv8-assignments/assignment2/httpserver/repositories"
	"github.com/rifqoi/hacktiv8-assignments/assignment2/httpserver/repositories/models"
	"gorm.io/gorm"
)

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) repositories.OrderRepo {
	return &orderRepo{
		db: db,
	}
}

func (o *orderRepo) GetOrders(limit int) (*[]models.Order, error) {
	var orders []models.Order
	results := o.db.Limit(limit).Preload("Items").Find(&orders)
	if results.Error != nil {
		return nil, results.Error
	}
	if results.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &orders, nil
}

func (o *orderRepo) CreateOrder(order *models.Order) error {
	error := o.db.Create(order).Error
	return error
}

func (o *orderRepo) UpdateOrder(order *models.Order) error {
	err := o.db.Debug().Model(&order).Association("Items").Replace(order.Items)
	if err != nil {
		return err
	}
	err = o.db.Debug().Session(&gorm.Session{FullSaveAssociations: true}).Preload("Items").Save(&order).Error
	return err
}

func (o *orderRepo) DeleteOrderByID(id int, order *models.Order) error {
	res := o.db.Where("id", id).Find(order)
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	err := o.db.Where("id", id).Delete(order).Error
	return err
}
