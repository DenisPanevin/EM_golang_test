package Helpers

import (
	"github.com/go-playground/validator/v10"
	"time"
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
	v.validator.RegisterValidation("date", v.ValidateDate)
	v.validator.RegisterValidation("settings", v.ValidateSettings)

	return &v
}

func (v *Validator) Validate(input interface{}) error {

	err := v.validator.Struct(input)
	if err != nil {
		println(err.Error())
		return err
	}

	/*for _, err := range err.(validator.ValidationErrors) {

		fmt.Println(err.Field())
		fmt.Println(err.Tag())
	}*/

	return nil
}

func (v *Validator) ValidateDate(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()
	_, err := time.Parse("2006-01-02", dateStr)
	return err == nil
}

func (v *Validator) ValidateSettings(fl validator.FieldLevel) bool {
	input, ok := fl.Field().Interface().([]int)
	if !ok {
		return false
	}

	output := true

	for i := 0; i < len(input); i++ {
		switch i {
		case 0:
			output = (input[0] < 365 && input[0] >= 0)

		case 1:
			output = input[1] >= 0

		}
	}

	return output

}
