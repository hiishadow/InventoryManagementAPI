package inventory

type InventoryService interface {
	CreateItem(req CreateItem) (Item, error)
	GetItemByID(id string) (Item, error)
	UpdateItemByID(id string, ) (Item, error)
	DeleteItemByID(id string) error
	GetAllItemsInLastestMonthByProductName(productName string) ([]Item, error)
	GetProductLastestMonthDataByProductName(productName string) ([]Item,ProductData , error)
}

type ItemPostgresRepository interface {
	Create(creatingItem CreateItem) (Item, error)
	GetByID(id string) (Item, error)
	UpdateByID(id string, updatingItem Item) (Item, error)
	DeleteByID(id string) error
	GetAllInLastestMonthByProductName(productName string) ([]Item, error)
}
