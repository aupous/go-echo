package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/aupous/go-echo/internal/app/go-echo/models"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// CreateUser create new user
func CreateUser(c echo.Context) error {
	u := new(models.User)
	c.Bind(&u)

	tx := c.Get("Tx").(*sqlx.Tx)

	user := models.NewUser(u.Name, u.Email, u.Password)

	if err := user.Create(tx); err != nil {
		logrus.Debug(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, user)
}

// GetUsers return all users
func GetUsers(c echo.Context) error {
	tx := c.Get("Tx").(*sqlx.Tx)
	users := new(models.Users)
	if err := users.FindAll(tx); err != nil {
		logrus.Debug(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, users)
}

// GetUser find user by id and return
func GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// handle error
		fmt.Println(err)
		return err
	}
	tx := c.Get("Tx").(*sqlx.Tx)
	user := new(models.User)

	if getErr := user.FindByID(tx, id); getErr != nil {
		logrus.Debug(err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, user)
}
