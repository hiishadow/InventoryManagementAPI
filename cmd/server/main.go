package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hiishadow/InventoryManagementAPI/pkg/config"
	postgresStore"github.com/hiishadow/InventoryManagementAPI/pkg/store/postgres"
	"github.com/hiishadow/InventoryManagementAPI/pkg/util"
)

func main() {
	config.LoadAppConfig()

	app := fiber.New(config.FiberConfig())

	// Init stores
	postgresStore.ConnectDB()

	util.SigHandler(app)
}
