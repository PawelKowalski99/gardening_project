package handlers

import (
	"net/http"
	// "encoding/json"
	"log"

	"github.com/PawelKowalski99/gardener_project/backend/db"
	"github.com/PawelKowalski99/gardener_project/backend/models"

	// "github.com/google/uuid"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

//----------
// Handlers
//----------

var (
	
	user = models.User{}
)

func CreateUser(c echo.Context) error {
	db := db.Manager()

	u := models.User{
		Name: 	c.FormValue("name"),
	}
	db.Create(&u)
	if err := c.Bind(&u); err != nil {
		return err
	}
	// users[u.ID] = u
	return c.JSON(http.StatusCreated, &u)
}


func GetAllUsers(db *gorm.DB) func(c echo.Context) error {
	return func(c echo.Context) error {
		var users []models.User
		db.Find(&users)
		log.Println(users)

		return c.JSON(http.StatusOK, users)
	}
}

func GetUser(c echo.Context) error {
	db := db.Manager()
	users := []models.User{}
	id := c.Param("id")
	u := db.Where("ID = ?", id).Find(&users)
	log.Println("1")
	log.Println(u)
	return c.JSON(http.StatusOK, u)
}

func UpdateUser(c echo.Context) error {
	u := models.User{}
	if err := c.Bind(u); err != nil {
		return err
	}
	// id := c.Param("id")
	// users[id].Name = u.Name
	return c.JSON(http.StatusOK, u)
}

func DeleteUser(c echo.Context) error {
	db := db.Manager()
	id:= c.Param("id")
	db.Delete(&models.User{}, id)
	return c.NoContent(http.StatusNoContent)
}