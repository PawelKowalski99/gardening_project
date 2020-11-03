package db

import (
	"github.com/PawelKowalski99/gardener_project/backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "fmt"
)

var db *gorm.DB 
var err error

//Init the database connection
func Init() {

	dsn := "user=gardening_project_db password=gardening_project_db dbname=go_restapi port=5432 sslmode=disable"
	
	db, err = gorm.Open(postgres.New(postgres.Config{
						DSN: dsn,
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
	db.AutoMigrate(&models.User{})
}

// Manager is a pointer to db
func Manager() *gorm.DB {
	return db
}