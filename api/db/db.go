package db

import (
	"os"

	"github.com/PawelKowalski99/gardener_project/backend/api/models"
	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "fmt"
)

//ConnectDB inits the database connection
func ConnectDB() (*gorm.DB, error) {
	err := godotenv.Load()

	dsn := os.Getenv("DB_DSN")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	},
	),
		&gorm.Config{
			PrepareStmt: true,
		},
	)
	if err != nil {
		panic("DB connection Error")
	}

	db.Migrator().DropTable(&models.User{}, &models.Subscription{}, &models.Order{})
	db.AutoMigrate(&models.User{}, &models.Subscription{}, &models.Order{})

	return db, nil
}
