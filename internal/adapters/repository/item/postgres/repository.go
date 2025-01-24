package repository

import (
	"errors"
	"time"

	"github.com/hiishadow/InventoryManagementAPI/internal/adapters/model"
	"github.com/hiishadow/InventoryManagementAPI/internal/core/inventory"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

// type ItemPostgresRepository interface {
// 	Create(creatingItem CreateItem) (Item, error)
// 	GetByID(id string) (Item, error)
// 	UpdateByID(id string, updatingItem Item) (Item, error)
// 	DeleteByID(id string) error
// 	GetAllInLastestMonthByProductName(productName string) ([]Item, error)
// 	GetAllBeforeDateByProductName(productName string, date time.Time) ([]Item, error)
// }

func NewItemPostgresRepository(db *gorm.DB) *repository {
	return &repository{DB: db}
}

func (r *repository) Create(creatingItem inventory.CreateItem) (inventory.Item, error) {
	var itemModel model.Item

	if err := copier.Copy(&itemModel, &creatingItem); err != nil {
		return inventory.Item{}, err
	}

	if err := r.DB.Create(&itemModel).Error; err != nil {
		return inventory.Item{}, err
	}

	var createdItem inventory.Item
	if err := copier.Copy(&createdItem, &itemModel); err != nil {
		return inventory.Item{}, err
	}

	return createdItem, nil
}

func (r *repository) GetByID(id string) (inventory.Item, error) {
	var itemModel model.Item
	if err := r.DB.Where("id = ?", id).First(&itemModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return inventory.Item{}, errors.New("Item not found")
		}
		return inventory.Item{}, err
	}

	var getItem inventory.Item
	if err := copier.Copy(&getItem, &itemModel); err != nil {
		return inventory.Item{}, err
	}

	return getItem, nil
}

func (r *repository) UpdateByID(id string, updatingItem inventory.UpdateItem) (inventory.Item, error) {
	var itemModel model.Item
	if err := r.DB.Where("id = ?", id).First(&itemModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return inventory.Item{}, errors.New("Item not found")
		}
		return inventory.Item{}, err
	}

	if err := copier.Copy(&itemModel, &updatingItem); err != nil {
		return inventory.Item{}, err
	}

	if err := r.DB.Model(&itemModel).Where("id = ?", id).Updates(itemModel).Error; err != nil {
		return inventory.Item{}, err
	}

	var updatedItem inventory.Item
	if err := copier.Copy(&updatedItem, &itemModel); err != nil {
		return inventory.Item{}, err
	}

	return updatedItem, nil
}

func (r *repository) DeleteByID(id string) error {
	if err := r.DB.Where("id = ?", id).Delete(&model.Item{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("Item not found")
		}
		return err
	}

	return nil
}

func (r *repository) GetAllInLastestMonthByProductName(productName string) ([]inventory.Item, error) {
	var itemsModel []model.Item
	now := time.Now()
	oneMonthAgo := now.AddDate(0, -1, 0)

	if err := r.DB.Where("product_name = ? AND at >= ?", productName, oneMonthAgo).Find(&itemsModel).Error; err != nil {
		return nil, err
	}

	var items []inventory.Item
	if err := copier.Copy(&items, &itemsModel); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *repository) GetAllBeforeDateByProductName(productName string, date time.Time) ([]inventory.Item, error) {
	var itemsModel []model.Item

	if err := r.DB.Where("product_name = ? AND at < ?", productName, date).Find(&itemsModel).Error; err != nil {
		return nil, err
	}

	var items []inventory.Item
	if err := copier.Copy(&items, &itemsModel); err != nil {
		return nil, err
	}

	return items, nil
}
