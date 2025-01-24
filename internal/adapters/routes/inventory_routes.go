package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hiishadow/InventoryManagementAPI/internal/adapters/handler"
)

func NewInventoryRoutes(route fiber.Router, handler *handler.InventoryHTTPHandler) {
	inventoryRoute := route.Group("")
	inventoryRoute.Post("/items", handler.CreateItem)
	inventoryRoute.Get("/items/:id", handler.GetItemByID)
	inventoryRoute.Patch("/items/:id", handler.UpdateItemByID)
	inventoryRoute.Delete("/items/:id", handler.DeleteItemByID)
	inventoryRoute.Get("/:productName", handler.GetProductLastestMonthDataByProductName)
}
