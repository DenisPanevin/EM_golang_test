package Helpers

import (
	"github.com/go-playground/validator/v10"
	"strconv"
	"strings"
)

type Ivalidator interface {
	Validate(input interface{}) error
}

type Validator struct {
	validator *validator.Validate
}

func NewValidator() Ivalidator {
	v := Validator{
		validator: validator.New(),
	}
	v.validator.RegisterValidation("passportNumber", v.PassportNumber)

	v.validator.RegisterValidation("checkFilterId", v.CheckFilterId)

	return &v
}

func (v *Validator) Validate(input interface{}) error {

	err := v.validator.Struct(input)
	if err != nil {
		println(err.Error())
		return err
	}

	return nil
}

func (v *Validator) PassportNumber(fl validator.FieldLevel) bool {
	passportStr := fl.Field().String()
	parts := strings.Split(passportStr, " ")
	if len(parts) != 2 {
		return false
	}
	if len(parts[0]) != 4 || len(parts[1]) != 6 {
		return false
	}
	number, err := strconv.Atoi(parts[0])
	if err != nil || number < 0 {
		return false
	}
	serie, err := strconv.Atoi(parts[1])
	if err != nil || serie < 0 {
		return false
	}
	return true
}

func (v *Validator) CheckFilterId(fl validator.FieldLevel) bool {
	id := fl.Field().Int()
	if id <= 0 {
		return false
	}

	return true
}
