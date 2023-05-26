package database

import (
	"log"
	"os"

	_ "database/sql"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func InitDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	dburi := os.Getenv("DB_URI") //used a cockroachdb database but postgres is fine
	DB, err = gorm.Open("postgres", dburi)
	if err != nil {
		log.Fatal("failed to connect database", err)
	} else {
		log.Println("connected to database")
	}
}
