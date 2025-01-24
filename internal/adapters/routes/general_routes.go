package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hiishadow/InventoryManagementAPI/internal/adapters/handler"
)

func NewGeneralRoutes(route fiber.Router, handler *handler.GeneralHTTPHandler) {
	route.Get("/", handler.Root)
	route.Use(handler.NotFound)
}
