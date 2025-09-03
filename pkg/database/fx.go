package database

import (
	"log"

	"upnext-fng/tgpl-api/config"
)

// Provider represents the function to provide dependency to fx modules
func Provider(cfg config.Config) Connection {
	defaultConfig := newDefaultConfig()

	// Use pkg logger for GORM logging
	defaultConfig.Logger = configureGormLogger(cfg)

	databaseConnection := NewConnection(cfg.Database.DataSourceName(), defaultConfig)
	err := databaseConnection.Open()
	if err != nil {
		log.Fatal("legal counsel database error:", err)
	}

	err = databaseConnection.Ping()
	if err != nil {
		log.Fatal("legal counsel database error:", err)
	}

	return databaseConnection
}
