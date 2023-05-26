package models

import "main/models/database"

// migrate models
func Migrate() {
	database.DB.AutoMigrate(&User{})
}
