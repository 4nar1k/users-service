package database

import (
	"errors"
	"github.com/4nar1k/users-service/internal/user"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func InitDB() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		logrus.Error("DATABASE_URL is not set")
		return nil, errors.New("DATABASE_URL is not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.WithError(err).Error("Failed to connect to database")
		return nil, err
	}

	// Автомиграция модели User
	if err := db.AutoMigrate(&user.User{}); err != nil {
		logrus.WithError(err).Error("Failed to auto-migrate User model")
		return nil, err
	}

	logrus.Info("Successfully connected to database")
	return db, nil
}
