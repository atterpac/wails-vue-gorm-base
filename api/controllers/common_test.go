package controllers

import (
	"changeme/api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Initialize a test db and app instance
func initApp() *App {
	Db, err := gorm.Open(sqlite.Open("./testData/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = Db.AutoMigrate(&models.User{})
	if err != nil {
		panic("failed to migrate database")
	}
	return &App{Db: Db}
}

