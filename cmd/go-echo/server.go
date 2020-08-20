package main

import (
	"github.com/aupous/go-echo/internal/app/go-echo/route"
	"github.com/go-playground/validator/v10"
)

type (
	// User type
	User struct {
		Name  string `json:"name" validate:"required"`
		Email string `json:"email" validate:"required,email"`
	}

	// CustomValidator for validate
	CustomValidator struct {
		validator *validator.Validate
	}
)

// Validate validate CustomValidator
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	router := route.Init()
	// e.Validator = &CustomValidator{validator: validator.New()}
	// e.POST("/users", func(c echo.Context) (err error) {
	// 	u := new(User)
	// 	if err = c.Bind(u); err != nil {
	// 		return
	// 	}
	// 	if err = c.Validate(u); err != nil {
	// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	// 	}
	// 	return c.JSON(http.StatusOK, u)
	// })
	router.Logger.Fatal(router.Start(":1323"))
}
