package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/nafeem-evatix/echoexplore/middlewares"
	"github.com/nafeem-evatix/echoexplore/v1/todos"
)

func main() {
	// Echo instance
	e := echo.New()

	// Global Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middlewares.Custom)

	// Grouping version 1 of api
	v1 := e.Group("api/v1/")

	initializeV1Services(v1)

	runAndGracefulShutdown(e)
}

func initializeV1Services(group *echo.Group) {
	// initializing todos service
	todos.Initialize(group)
}

func runAndGracefulShutdown(e *echo.Echo) {
	// Start server
	go func() {
		if err := e.Start(":1323"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
