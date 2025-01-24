package inventory

import (
	"github.com/jinzhu/copier"
)

// type InventoryService interface {
// 	CreateItem(req CreateItem) (Item, error)
// 	GetItemByID(id string) (Item, error)
// 	UpdateItemByID(id string, ) (Item, error)
// 	DeleteItemByID(id string) error
// 	GetAllItemsInLastestMonthByProductName(productName string) ([]Item, error)
// 	GetProductLastestMonthDataByProductName(productName string) ([]Item,ProductData , error)
// }

type service struct {
	ItemPostgresRepository ItemPostgresRepository
}

func NewInventoryService(repo ItemPostgresRepository) *service {
	return &service{ItemPostgresRepository: repo}
}

// func (s *service) CreateItem(creatingItem CreateItem) (Item, error) {
// 	return s.ItemPostgresRepository.Create(creatingItem)
// }

// func (s *service) GetItemByID(id string) (GetItem, error) {
// }

func (s *service) UpdateItemByID(id string, updatingItem Item) (Item, error) {
	return s.ItemPostgresRepository.UpdateByID(id, updatingItem)
}

func (s *service) DeleteItemByID(id string) error {
	return s.ItemPostgresRepository.DeleteByID(id)
}

type ItemInventory struct {
	totalQuantity int
	totalCost     float64
	averageCost   float64
	pnlResult     float64
}

func calculatePNL(items []Item) ItemInventory {
	itemInventory := ItemInventory{}

	// Iterate through the transactions
	for _, item := range items {
		if item.Status == "BUY" {
			// Update inventory based on the BUY transaction
			itemInventory.totalQuantity += item.Amount
			itemInventory.totalCost += float64(item.Amount) * item.Price

			// Recalculate the average cost
			if itemInventory.totalQuantity > 0 {
				itemInventory.averageCost = itemInventory.totalCost / float64(itemInventory.totalQuantity)
			}
		} else if item.Status == "SELL" {
			// Calculate PNL based on the current average cost
			cogs := float64(item.Amount) * itemInventory.averageCost
			revenue := float64(item.Amount) * item.Price
			itemInventory.pnlResult += (revenue - cogs)

			// Update inventory after the sale
			itemInventory.totalQuantity -= item.Amount
			itemInventory.totalCost -= cogs

			// Avoid division by zero when inventory becomes empty
			if itemInventory.totalQuantity == 0 {
				itemInventory.averageCost = 0
			}
		}
	}

	return itemInventory
}

func addcalculatedPNLToItemsAndGetProductData(items []Item) ([]GetItem, ProductData) {
	addedPNLItems := []GetItem{}
	itemInventory := ItemInventory{}
	productData := ProductData{}

	// Iterate through the transactions
	for _, item := range items {
		if item.Status == "BUY" {
			// Update inventory based on the BUY transaction
			itemInventory.totalQuantity += item.Amount
			itemInventory.totalCost += float64(item.Amount) * item.Price
			productData.ProductsBought += item.Amount

			// Recalculate the average cost
			if itemInventory.totalQuantity > 0 {
				itemInventory.averageCost = itemInventory.totalCost / float64(itemInventory.totalQuantity)
			}

			var addedPNLItem GetItem
			copier.Copy(&addedPNLItem, &item)
			addedPNLItems = append(addedPNLItems, addedPNLItem)

		} else if item.Status == "SELL" {
			// Calculate PNL based on the current average cost
			cogs := float64(item.Amount) * itemInventory.averageCost
			revenue := float64(item.Amount) * item.Price
			pnl := (revenue - cogs)
			itemInventory.pnlResult += pnl

			var addedPNLItem GetItem
			copier.Copy(&addedPNLItem, &item)
			addedPNLItem.PNL = pnl
			addedPNLItems = append(addedPNLItems, addedPNLItem)

			// Update inventory after the sale
			itemInventory.totalQuantity -= item.Amount
			itemInventory.totalCost -= cogs
			productData.ProductsSold += item.Amount

			// Avoid division by zero when inventory becomes empty
			if itemInventory.totalQuantity == 0 {
				itemInventory.averageCost = 0
			}
		}
	}
	productData.TotalAmount = itemInventory.totalQuantity
	productData.Profit = itemInventory.pnlResult
	return addedPNLItems, productData
}
