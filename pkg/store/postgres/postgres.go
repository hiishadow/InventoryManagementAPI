package postgres

import "gorm.io/gorm"

var DB *gorm.DB

func GetDBClient() *gorm.DB {
	return DB
}
