package handlers

import ( // "encoding/json"
	// "github.com/google/uuid"

	"log"
	"net/http"
	"strconv"
	"time"

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
	price, _ := strconv.ParseFloat(c.FormValue("price"), 64)

	months, err := strconv.Atoi(c.FormValue("months"))
	if err != nil {
		log.Println(months)
		return err
	}

	s := models.Subscription{
		UserID:      userID,
		Price:       price,
		Description: c.FormValue("description"),
		TimeEnd:     time.Now().AddDate(0, months, 0),
	}
	h.db.Create(&s)
	return c.JSONPretty(http.StatusOK, s, " ")

}

func (h *Handler) GetOrder(c echo.Context) error {
	//Auth jwt token
	u := c.Get("user")
	claims := u.(*jwt.Token).Claims.(jwt.MapClaims)
	userID := uint(claims["id"].(float64))

	var subscriptions []models.Subscription
	s := models.Subscription{}
	h.db.Where("user_id  = ?", userID).Find(&subscriptions).Scan(&s)
	log.Println(s)
	return c.JSONPretty(http.StatusOK, s, " ")
}

func (h *Handler) UpdateOrder(c echo.Context) error {
	//Auth jwt token
	u := c.Get("user")
	claims := u.(*jwt.Token).Claims.(jwt.MapClaims)
	userID := uint(claims["id"].(float64))

	var subscription models.Subscription
	var subscriptions []models.Subscription
	h.db.Where("user_id  = ?", userID).Find(&subscriptions).Scan(&subscription)

	update := models.Subscription{}
	if err := c.Bind(&update); err != nil {
		return err
	}
	h.db.Model(&subscription).Updates(update)
	h.db.Save(&subscription)

	return c.JSONPretty(http.StatusOK, subscription, " ")
}

func (h *Handler) DeleteOrder(c echo.Context) error {
	//Auth jwt token
	u := c.Get("user")
	claims := u.(*jwt.Token).Claims.(jwt.MapClaims)
	userID := uint(claims["id"].(float64))

	h.db.Delete(&models.Subscription{}, userID)
	return c.NoContent(http.StatusNoContent)
}
