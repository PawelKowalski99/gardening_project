package handlers

import (
	"net/http"

	// "encoding/json"
	"log"
	"strconv"

	"github.com/PawelKowalski99/gardener_project/backend/api/models"

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

//Handler holds everythinng, that controller needsd
type Handler struct {
	db *gorm.DB
}

//NewHandler returns a new BaseHandler
func NewHandler(db *gorm.DB) *Handler {
	return &Handler{
		db: db,
	}
}

//CreateUser creates user
func (h *Handler) CreateUser(c echo.Context) error {
	jobRange, _ := strconv.ParseFloat(c.FormValue("range"), 64)

	u := models.User{
		FirstName:    c.FormValue("first_name"),
		LastName:     c.FormValue("last_name"),
		Email:        c.FormValue("email"),
		Password:     c.FormValue("password"),
		Order:        models.Order{},
		Subscription: models.Subscription{},
		Role:         c.FormValue("role"),
		Range:        jobRange,
	}

	h.db.Create(&u)
	return c.JSONPretty(http.StatusCreated, &u, " ")
}

//GetAllUsers gets all users from table User
func (h *Handler) GetAllUsers(c echo.Context) error {
	var users []models.User
	h.db.Find(&users)
	log.Println(users)

	return c.JSON(http.StatusOK, users)

}

func (h *Handler) GetUser(c echo.Context) error {
	users := []models.User{}
	id := c.Param("id")
	u := h.db.Where("ID = ?", id).Find(&users)
	log.Println(u)
	return c.JSON(http.StatusOK, u)
}

func (h *Handler) UpdateUser(c echo.Context) error {
	id := c.FormValue("id")
	u := models.User{}
	update := models.User{}
	h.db.Find(&u, id)
	if err := c.Bind(&update); err != nil {
		return err
	}
	log.Println(update)
	h.db.Model(&u).Updates(update)
	h.db.Save(&u)
	return c.JSON(http.StatusOK, u)
}

func (h *Handler) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	h.db.Delete(&models.User{}, id)
	return c.NoContent(http.StatusNoContent)
}
