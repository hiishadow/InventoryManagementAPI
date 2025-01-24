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

func CreateType() {
	err := DB.Exec(`
	DO $$ BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'item_status') THEN
					CREATE TYPE item_status AS ENUM ('BUY', 'SELL');
			END IF;
	END $$;
`).Error
	if err != nil {
		log.Fatalf("Can't create ENUM type. error: %v", err)
		return
	}
}
