package models

import "main/models/database"

func Migrate() {
	database.DB.AutoMigrate(&User{})
}
