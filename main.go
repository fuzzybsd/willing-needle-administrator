package main

import (
	"net/http"

	"github.com/fuzzybsd/willing_needle-administrator/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Debug = true

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	h := handler.NewUserHandler()
	e.POST("/users", h.CreateUser)
	e.GET("/users/:name", h.GetUser)
	e.PUT("/users/:name", h.UpdateUser)
	e.DELETE("/users/:name", h.DeleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}
