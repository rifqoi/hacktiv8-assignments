package gorm

import (
	"github.com/rifqoi/hacktiv8-assignments/assignment2/httpserver/repositories"
	"github.com/rifqoi/hacktiv8-assignments/assignment2/httpserver/repositories/models"
	"gorm.io/gorm"
)

type itemRepo struct {
	db *gorm.DB
}

func NewItemRepo(db *gorm.DB) repositories.ItemRepo {
	return &itemRepo{
		db: db,
	}
}

func (i *itemRepo) CreateItem(item *models.Item) error {
	err := i.db.Create(item).Error
	return err
}
