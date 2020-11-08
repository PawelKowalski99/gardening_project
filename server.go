package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/PawelKowalski99/gardener_project/backend/api"

	// "fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	e := echo.New()

	err := api.SetRouters(e)
	if err != nil {

	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	go func() {
		if err := e.Start(":1323"); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}
