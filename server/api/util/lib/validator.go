package lib

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func NewValidator() (*Validator, error) {
	validate := validator.New()
	return &Validator{
		validator: validate,
	}, nil
}

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	return false
}


