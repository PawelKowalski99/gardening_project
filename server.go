package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/PawelKowalski99/gardener_project/backend/api"
	"github.com/PawelKowalski99/gardener_project/backend/api/db"
	"gorm.io/gorm"

	// "fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type App struct {
	e  *echo.Echo
	db *gorm.DB
}

func (a *App) Initialize() error {
	a.e = echo.New()

	var err error
	a.db, err = db.ConnectDB()
	if err != nil {
		return err
	}

	err = api.SetRouters(a.e, a.db)
	if err != nil {

	}

	// Middleware
	a.e.Use(middleware.Logger())
	a.e.Use(middleware.Recover())

	return nil
}

func (a *App) Run() error {
	go func() {
		if err := a.e.Start(":1323"); err != nil {
			a.e.Logger.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := a.e.Shutdown(ctx); err != nil {
		a.e.Logger.Fatal(err)
	}
	return nil
}

func main() {
	app := App{}

	app.Initialize()
	app.Run()

}
