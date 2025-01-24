package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hiishadow/InventoryManagementAPI/internal/adapters/model"
	"github.com/hiishadow/InventoryManagementAPI/pkg/config"
	postgresStore "github.com/hiishadow/InventoryManagementAPI/pkg/store/postgres"
	"github.com/hiishadow/InventoryManagementAPI/pkg/util"
)

func main() {
	config.LoadAllconfig()

	app := fiber.New(config.FiberConfig())


	// Init stores
	postgresStore.ConnectDB()
	postgresStore.CreateType()
	postgresStore.MigrateDB(
		&model.Item{},
	)

	util.SigHandler(app)
}
