package api

import (
	"github.com/PawelKowalski99/gardener_project/backend/api/db"
	"github.com/PawelKowalski99/gardener_project/backend/api/handlers"
	"github.com/labstack/echo"
)

func SetRouters(e *echo.Echo) error {
	db, err := db.ConnectDB()
	if err != nil {
		return err
	}
	h := handlers.NewHandler(db)

	userGroup := e.Group("/users")

	userGroupFunc(userGroup, h)

	return nil
}

func userGroupFunc(g *echo.Group, h *handlers.Handler) {
	g.GET("", h.GetAllUsers)
	g.POST("", h.CreateUser)
	g.GET("/:id", h.GetUser)
	g.PUT("/:id", h.UpdateUser)
	g.DELETE("/:id", h.DeleteUser)
}

// func AuthGroup
