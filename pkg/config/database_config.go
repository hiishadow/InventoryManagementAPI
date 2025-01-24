package config

import (
	"os"
)

type DB struct {
	Host     string
	Name     string
	Username string
	Password string
	Port     string
	SslMode  string
}

var db = &DB{}

func DBConfig() *DB {
	return db
}

func LoadDBConfig() {
	db.Host = os.Getenv("DB_HOST")
	db.Name = os.Getenv("DB_NAME")
	db.Username = os.Getenv("DB_USERNAME")
	db.Password = os.Getenv("DB_PASSWORD")
	db.Port = os.Getenv("DB_PORT")
	db.SslMode = os.Getenv("DB_SSL_MODE")
}
