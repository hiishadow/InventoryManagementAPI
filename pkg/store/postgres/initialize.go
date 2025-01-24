package postgres

import (
	"fmt"

	"github.com/hiishadow/InventoryManagementAPI/pkg/config"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error

	dbConfig := config.DBConfig()

	DB, err = gorm.Open(postgres.Open(
		fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
			dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name, dbConfig.SslMode),
	), &gorm.Config{})
	if err != nil {
		log.Fatalf("Can't connect to database. error: %v", err)
		return
	}
	log.Info("Database connected")
}

func MigrateDB(models ...interface{}) {	
	for _, model := range models {
		if err := DB.AutoMigrate(model); err != nil {
			log.Fatalf("Can't migrate database. error: %v", err)
			return
		}
	}

	log.Info("Database migrated")
}
