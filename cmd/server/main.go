package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hiishadow/InventoryManagementAPI/internal/adapters/handler"
	"github.com/hiishadow/InventoryManagementAPI/internal/adapters/model"
	itemPostgresRepository "github.com/hiishadow/InventoryManagementAPI/internal/adapters/repository/item/postgres"
	"github.com/hiishadow/InventoryManagementAPI/internal/adapters/routes"
	inventoryService "github.com/hiishadow/InventoryManagementAPI/internal/core/inventory"
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

	// Init repositories
	itemPostgresRepository := itemPostgresRepository.NewItemPostgresRepository(postgresStore.DB)

	// Init services
	inventoryService := inventoryService.NewInventoryService(itemPostgresRepository)

	// Init handlers
	generalHTTPHandler := handler.NewGeneralHTTPHandler()
	inventoryHTTPHandler := handler.NewInventoryHTTPHandler(inventoryService)

	// Init routes
	inventoryRoute := app.Group("/inventory")
	routes.NewInventoryRoutes(inventoryRoute, inventoryHTTPHandler)
	routes.NewGeneralRoutes(app, generalHTTPHandler)

	util.SigHandler(app)
}
