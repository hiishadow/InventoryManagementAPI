package handler

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/hiishadow/InventoryManagementAPI/internal/core/inventory"
	"github.com/jinzhu/copier"
)

type InventoryHTTPHandler struct {
	InventoryService inventory.InventoryService
}

func NewInventoryHTTPHandler(inventoryService inventory.InventoryService) *InventoryHTTPHandler {
	return &InventoryHTTPHandler{InventoryService: inventoryService}
}

func (h *InventoryHTTPHandler) CreateItem(c *fiber.Ctx) error {
	var body inventory.CreateItemRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	parsedTime, err := time.Parse(time.RFC3339, body.At)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	createdItem, err := h.InventoryService.CreateItem(inventory.CreateItem{
		ProductName: body.ProductName,
		Status:      body.Status,
		Price:       body.Price,
		Amount:      body.Amount,
		At:          parsedTime,
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"id": createdItem.ID,
	})
}

func (h *InventoryHTTPHandler) GetItemByID(c *fiber.Ctx) error {
	id := c.Params("id")

	item, err := h.InventoryService.GetItemByID(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(item)
}

func (h *InventoryHTTPHandler) UpdateItemByID(c *fiber.Ctx) error {
	id := c.Params("id")

	var body inventory.UpdateItemRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	var updatingItem inventory.UpdateItem
	if err := copier.Copy(&updatingItem, &body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if body.At != "" {
		parsedTime, err := time.Parse(time.RFC3339, body.At)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		updatingItem.At = parsedTime
	}
	updatedItem, err := h.InventoryService.UpdateItemByID(id, updatingItem)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"id": updatedItem.ID,
	})
}

func (h *InventoryHTTPHandler) DeleteItemByID(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.InventoryService.DeleteItemByID(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"id": id,
	})
}

func (h *InventoryHTTPHandler) GetProductLastestMonthDataByProductName(c *fiber.Ctx) error {
	productName := c.Params("productName")

	items, productData, err := h.InventoryService.GetProductLastestMonthDataByProductName(productName)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data":                        items,
		"totalAmount":                 productData.TotalAmount,
		"productsSoldInLatestMonth":   productData.ProductsSold,
		"productsBoughtInLatestMonth": productData.ProductsBought,
		"latestMonthProfit":           productData.Profit,
	})
}
