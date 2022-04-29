package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/fuzzybsd/willing_needle-core/model"
)

type (
	UserHandler struct {
		db map[string]*model.User
	}
)

func NewUserHandler() *UserHandler {
	h := new(UserHandler)
	m := make(map[string]*model.User)
	m["hoge"] = &model.User{"hoge", "hoge@email"}
	m["foo"] = &model.User{"foo", "foo@email"}
	h.db = m
	return h
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}

func (h *UserHandler) GetUser(c echo.Context) error {
	c.Logger().Print("ghohohgoe")
	email := c.Param("email")
	user := h.db[email]
	if user == nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}
