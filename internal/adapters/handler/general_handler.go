package handler

import "github.com/gofiber/fiber/v2"

type GeneralHTTPHandler struct{}

func NewGeneralHTTPHandler() *GeneralHTTPHandler {
	return &GeneralHTTPHandler{}
}

func (h *GeneralHTTPHandler) Root(c *fiber.Ctx) error {
	return c.SendString("Inventory Management API is running")
}

func (h *GeneralHTTPHandler) NotFound(c *fiber.Ctx) error {
	return c.SendString("Not Found")
}
