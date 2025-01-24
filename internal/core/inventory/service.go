package inventory

import (
	"errors"

	"github.com/jinzhu/copier"
)

type service struct {
	ItemPostgresRepository ItemPostgresRepository
}

func NewInventoryService(repo ItemPostgresRepository) *service {
	return &service{ItemPostgresRepository: repo}
}

func (s *service) CreateItem(creatingItem CreateItem) (Item, error) {
	if creatingItem.Status == "SELL" {
		items, err := s.ItemPostgresRepository.GetAllBeforeDateByProductName(creatingItem.ProductName, creatingItem.At)
		if err != nil {
			return Item{}, err
		}

		itemInventory := calculatePNL(items)
		if creatingItem.Amount > itemInventory.totalQuantity {
			return Item{}, errors.New("Not enough stock to sell")
		}
	}
	return s.ItemPostgresRepository.Create(creatingItem)
}

func (s *service) GetItemByID(id string) (GetItem, error) {
	item, err := s.ItemPostgresRepository.GetByID(id)
	if err != nil {
		return GetItem{}, err
	}
	var addedPNLItem GetItem
	if item.Status == "SELL" {
		items, err := s.ItemPostgresRepository.GetAllBeforeDateByProductName(item.ProductName, item.At)
		if err != nil {
			return GetItem{}, err
		}

		itemInventory := calculatePNL(items)
		addedPNLItem.PNL = (float64(item.Amount) * item.Price) - (float64(item.Amount) * itemInventory.averageCost)
	}

	if err := copier.Copy(&addedPNLItem, &item); err != nil {
		return GetItem{}, err
	}

	return addedPNLItem, nil
}

func (s *service) UpdateItemByID(id string, updatingItem UpdateItem) (Item, error) {
	return s.ItemPostgresRepository.UpdateByID(id, updatingItem)
}

func (s *service) DeleteItemByID(id string) error {
	return s.ItemPostgresRepository.DeleteByID(id)
}

func (s *service) GetProductLastestMonthDataByProductName(productName string) ([]GetItem, ProductData, error) {
	items, err := s.ItemPostgresRepository.GetAllInLastestMonthByProductName(productName)
	if err != nil {
		return nil, ProductData{}, err
	}
	addedPNLItems, productData := addcalculatedPNLToItemsAndGetProductData(items)
	return addedPNLItems, productData, nil
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
