package util

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/hiishadow/InventoryManagementAPI/pkg/config"
	log "github.com/sirupsen/logrus"
)

func SigHandler(app *fiber.App) {
	appConfig := config.AppConfig()

	log.Infof("Server is started on http://%s:%d", appConfig.Host, appConfig.Port)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		<-sigCh
		log.Info("Shutting down server...")
		_ = app.Shutdown()
	}()

	serverAddr := fmt.Sprintf("%s:%d", appConfig.Host, appConfig.Port)
	if err := app.Listen(serverAddr); err != nil {
		log.Errorf("Oops... server is not running! error: %v", err)
	}
}
