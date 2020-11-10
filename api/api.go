package api

import (
	"github.com/PawelKowalski99/gardener_project/backend/api/db"
	"github.com/PawelKowalski99/gardener_project/backend/api/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetRouters(e *echo.Echo) error {
	db, err := db.ConnectDB()
	if err != nil {
		return err
	}
	h := handlers.NewHandler(db)

	userGroup := e.Group("/users")
	authGroup := e.Group("/auth")
	restrictedGroup := e.Group("/restricted")
	subscriptionGroup := restrictedGroup.Group("/subscriptions")
	orderGroup := restrictedGroup.Group("/orders")

	restrictedGroup.Use(middleware.JWT([]byte("secret")))

	userGroupMethods(userGroup, h)
	authGroupMethods(authGroup, h)
	subscriptionGroupMethods(subscriptionGroup, h)
	orderGroupMethods(orderGroup, h)

	return nil
}

func userGroupMethods(g *echo.Group, h *handlers.Handler) {
	g.GET("", h.GetAllUsers)
	g.POST("", h.CreateUser)
	g.GET("/:id", h.GetUser)
	g.PUT("/:id", h.UpdateUser)
	g.DELETE("/:id", h.DeleteUser)
}

func authGroupMethods(g *echo.Group, h *handlers.Handler) {
	g.POST("/login", h.Login)
}

func subscriptionGroupMethods(g *echo.Group, h *handlers.Handler) {
	// g.GET("", h.GetAllUsers)
	g.POST("", h.CreateSubscription)
	g.GET("/:id", h.GetSubscription)
	g.PUT("/:id", h.UpdateSubscription)
	g.DELETE("/:id", h.DeleteSubscription)
}

func orderGroupMethods(g *echo.Group, h *handlers.Handler) {
	// g.GET("", h.GetAllUsers)
	g.POST("", h.CreateOrder)
	g.GET("/:id", h.GetOrder)
	g.PUT("/:id", h.UpdateOrder)
	g.DELETE("/:id", h.DeleteOrder)
}
