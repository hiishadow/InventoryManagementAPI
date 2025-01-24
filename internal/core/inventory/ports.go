package inventory

import "time"

type InventoryService interface {
	CreateItem(creatingItem CreateItem) (Item, error)
	GetItemByID(id string) (GetItem, error)
	UpdateItemByID(id string, updatingItem UpdateItem) (Item, error)
	DeleteItemByID(id string) error
	// GetAllItemsInLastestMonthByProductName(productName string) ([]GetItem, error)
	GetProductLastestMonthDataByProductName(productName string) ([]GetItem, ProductData, error)
}

type ItemPostgresRepository interface {
	Create(creatingItem CreateItem) (Item, error)
	GetByID(id string) (Item, error)
	UpdateByID(id string, updatingItem UpdateItem) (Item, error)
	DeleteByID(id string) error
	GetAllInLastestMonthByProductName(productName string) ([]Item, error)
	GetAllBeforeDateByProductName(productName string, date time.Time) ([]Item, error)
}
