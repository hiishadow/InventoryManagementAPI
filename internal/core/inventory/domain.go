package inventory

import (
	"time"

	"github.com/google/uuid"
)

type ItemStatus string

const (
	BUY  ItemStatus = "BUY"
	SELL ItemStatus = "SELL"
)

type CreateItemRequest struct {
	ProductName string     `json:"productName" validate:"required"`
	Status      ItemStatus `json:"status" validate:"required,oneof=BUY SELL"`
	Price       float64    `json:"price" validate:"required"`
	Amount      int        `json:"amount" validate:"required"`
	At          string     `json:"at" validate:"required"`
}

type UpdateItemRequest struct {
	ProductName string     `json:"productName"`
	Status      ItemStatus `json:"status" validate:"oneof=BUY SELL"`
	Price       float64    `json:"price"`
	Amount      int        `json:"amount"`
	At          string     `json:"at"`
}

type CreateItem struct {
	ProductName string
	Status      ItemStatus
	Price       float64
	Amount      int
	At          time.Time
}

type Item struct {
	ID uuid.UUID
	CreateItem
}

type ProductData struct {
	TotalAmount    int
	ProductsSold   int
	ProductsBought int
	Profit         float64
}
