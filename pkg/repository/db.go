package repository

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strings"
	"test-task/config"
)

type DB struct {
	*gorm.DB
}

func Connection(c *config.Config) (*DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DBConf.Host, c.DBConf.Port, c.DBConf.Username, c.DBConf.Password, c.DBConf.DBName, c.DBConf.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
		Logger:                 logger.Default.LogMode(setLevel(c)),
	})

	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func setLevel(c *config.Config) logger.LogLevel {
	switch strings.ToLower(c.DBConf.LogLevel) {

	case "silent":
		return logger.Silent
	case "error":
		return logger.Error
	case "warn":
		return logger.Warn
	case "info":
		return logger.Info
	default:
		return logger.Info
	}

}
