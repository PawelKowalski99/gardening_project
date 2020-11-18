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

const (
	//Models
	user         = "user"
	subscription = "subscription"
	order        = "order"
)

//Handler holds everythinng, that controller needsd
type Handler struct {
	db *gorm.DB
}

type User struct {
	db *gorm.DB
}

type CRUDHandler interface {
	Create(c echo.Context) error
	Get(c echo.Context) error
	GetAll(c echo.Context) error
	Delete(c echo.Context) error
	Update(c echo.Context) error
}

func NewBaseHandler(db *gorm.DB) *Handler {
	return &Handler{db: db}
}

//NewHandler returns a new BaseHandler
func NewCRUDHandler(db *gorm.DB, name string) CRUDHandler {
	switch name {
	case user:
		return &User{db}
	case subscription:
		return &Subscription{db}
	case order:
		return &Order{db}
	default:
		return
	}
	return nil
}

//CreateUser creates user
func (u *User) Create(c echo.Context) error {
	jobRange, _ := strconv.ParseFloat(c.FormValue("range"), 64)

	user := models.User{
		FirstName:    c.FormValue("first_name"),
		LastName:     c.FormValue("last_name"),
		Email:        c.FormValue("email"),
		Password:     c.FormValue("password"),
		Order:        models.Order{},
		Subscription: models.Subscription{},
		Role:         c.FormValue("role"),
		Range:        jobRange,
	}

	u.db.Create(&user)
	return c.JSONPretty(http.StatusCreated, &user, " ")
}

//GetAllUsers gets all users from table User
func (u *User) GetAll(c echo.Context) error {
	var users []models.User
	u.db.Find(&users)
	log.Println(users)

	return c.JSON(http.StatusOK, users)

}

func (u *User) Get(c echo.Context) error {
	users := []models.User{}
	id := c.Param("id")
	user := models.User{}
	u.db.Where("ID = ?", id).Find(&users).Scan(&user)
	log.Println(u)
	return c.JSON(http.StatusOK, u)
}

func (u *User) Update(c echo.Context) error {
	id := c.FormValue("id")
	user := models.User{}
	update := models.User{}
	u.db.Find(&user, id)
	if err := c.Bind(&update); err != nil {
		return err
	}
	log.Println(update)
	u.db.Model(&user).Updates(update)
	u.db.Save(&user)
	return c.JSON(http.StatusOK, user)
}

func (h *User) Delete(c echo.Context) error {
	id := c.Param("id")
	h.db.Delete(&models.User{}, id)
	return c.NoContent(http.StatusNoContent)
}
