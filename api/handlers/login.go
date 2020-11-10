package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/PawelKowalski99/gardener_project/backend/api/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func (h *Handler) Login(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	users := []models.User{}
	user := models.User{}
	h.db.Where("email = ? AND password = ?", email, password).Find(&users).Scan(&user)

	// Throw unauthorized error
	if email != user.Email || password != user.Password {
		return echo.ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	b, _ := json.Marshal(user)

	return c.JSON(http.StatusOK, map[string]string{
		"token":   t,
		"id":      strconv.FormatUint(uint64(user.ID), 10),
		"user_id": strconv.FormatUint(uint64(claims["id"].(uint)), 10),
		"user":    string(b),
	})
}
