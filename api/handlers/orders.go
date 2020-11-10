package handlers

import ( // "encoding/json"
	// "github.com/google/uuid"

	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/PawelKowalski99/gardener_project/backend/api/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func (h *Handler) CreateOrder(c echo.Context) error {

	//Auth jwt token
	u := c.Get("user")
	claims := u.(*jwt.Token).Claims.(jwt.MapClaims)
	userID := uint(claims["id"].(float64))

	//Prepare struct fields
	diff, err := strconv.Atoi(c.FormValue("difficulty"))
	if diff < 0 || diff > 10 {
		return errors.New("Difficulty not in range (difficulty < 0 && difficulty > 10")
	} else if err != nil {
		return err
	}

	o := models.Order{
		UserID:      userID,
		Difficulty:  uint(diff),
		Description: c.FormValue("description"),
	}
	h.db.Create(&o)
	return c.JSONPretty(http.StatusOK, o, " ")
}

func (h *Handler) GetOrder(c echo.Context) error {
	//Auth jwt token
	u := c.Get("user")
	claims := u.(*jwt.Token).Claims.(jwt.MapClaims)
	userID := uint(claims["id"].(float64))

	var orders []models.Order
	o := models.Order{}
	h.db.Where("user_id  = ?", userID).Find(&orders).Scan(o)
	log.Println(o)
	return c.JSONPretty(http.StatusOK, o, " ")
}

func (h *Handler) UpdateOrder(c echo.Context) error {
	//Auth jwt token
	u := c.Get("user")
	claims := u.(*jwt.Token).Claims.(jwt.MapClaims)
	userID := uint(claims["id"].(float64))

	var orders []models.Order
	o := models.Order{}
	h.db.Where("user_id  = ?", userID).Find(&orders).Scan(o)

	update := models.Order{}
	if err := c.Bind(&update); err != nil {
		return err
	}
	h.db.Model(&o).Updates(update)
	h.db.Save(&o)

	return c.JSONPretty(http.StatusOK, o, " ")
}

func (h *Handler) DeleteOrder(c echo.Context) error {
	//Auth jwt token
	u := c.Get("user")
	claims := u.(*jwt.Token).Claims.(jwt.MapClaims)
	userID := uint(claims["id"].(float64))

	h.db.Delete(&models.Order{}, userID)
	return c.NoContent(http.StatusNoContent)
}
