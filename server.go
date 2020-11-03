package main

import (
	"github.com/PawelKowalski99/gardener_project/backend/db"
	"github.com/PawelKowalski99/gardener_project/backend/handlers"

	// "fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	db.Init()
	db := db.Manager()

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	// g := e.Group("/users")
	e.GET("/users", handlers.GetAllUsers(db))
	e.POST("/users", handlers.CreateUser)
	e.GET("/users/:id", handlers.GetUser)
	e.PUT("/users/:id", handlers.UpdateUser)
	e.DELETE("/users/:id", handlers.DeleteUser)

	// Start server

	e.Logger.Fatal(e.Start(":1323"))

}
