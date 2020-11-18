package api

import (
	"github.com/PawelKowalski99/gardener_project/backend/api/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/gorm"
)

type CRUDHandler interface {
	Create(c echo.Context) error
	Get(c echo.Context) error
	GetAll(c echo.Context) error
	Delete(c echo.Context) error
	Update(c echo.Context) error
}

type CRUDGroup struct {
	G       *echo.Group
	Handler CRUDHandler
}

func SetRouters(e *echo.Echo, db *gorm.DB) error {
	h := handlers.NewBaseHandler(db)

	authGroup := e.Group("/auth")
	restrictedGroup := e.Group("/restricted")

	restrictedGroup.Use(middleware.JWT([]byte("secret")))

	CRUDGroups := []CRUDGroup{
		{
			G:       e.Group("/users"),
			Handler: handlers.NewCRUDHandler(db, "user"),
		},
		{
			G:       restrictedGroup.Group("/subscriptions"),
			Handler: handlers.NewCRUDHandler(db, "subscription"),
		},
		{
			G:       restrictedGroup.Group("/orders"),
			Handler: handlers.NewCRUDHandler(db, "order"),
		},
	}

	for _, CRUDGroup := range CRUDGroups {
		groupCRUDMethods(CRUDGroup.G, CRUDGroup.Handler)
	}

	authGroupMethods(authGroup, h)

	return nil
}

func groupCRUDMethods(g *echo.Group, h CRUDHandler) {
	g.GET("", h.GetAll)
	g.POST("", h.Create)
	g.GET("/:id", h.Get)
	g.PUT("/:id", h.Update)
	g.DELETE("/:id", h.Delete)
}

func authGroupMethods(g *echo.Group, h *handlers.Handler) {
	g.POST("/login", h.Login)
}
