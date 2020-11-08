package db

import (
	"github.com/PawelKowalski99/gardener_project/backend/api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "fmt"
)

//ConnectDB inits the database connection
func ConnectDB() (*gorm.DB, error) {

	dsn := "user=gardening_project_db password=gardening_project_db dbname=gardening_project_db port=5432 sslmode=disable"

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
