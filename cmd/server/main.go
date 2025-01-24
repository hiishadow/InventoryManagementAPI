package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hiishadow/InventoryManagementAPI/pkg/config"
	"github.com/hiishadow/InventoryManagementAPI/pkg/util"
)

func main() {
	config.LoadAppConfig()

	app := fiber.New(config.FiberConfig())

	util.SigHandler(app)
}