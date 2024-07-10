package models

import (
	"encoding/json"
	"reflect"
	"time"
)

type CreateUserDto struct {
	PassportNumber string `json:"passportNumber,omitempty" valid:"required,minstringlength(5)"`
	Name           string `json:"name,omitempty" valid:"required,minstringlength(3)"`
	Surname        string `json:"surname,omitempty" valid:"required,minstringlength(3)"`
	Patronymic     string `json:"patronymic,omitempty" valid:"required,minstringlength(3)"`
	Address        string `json:"address,omitempty" valid:"required ,minstringlength(28)"`
}

type PassportNumberDto struct {
	PassportNumber string `json:"passportNumber" valid:"passportNumber,required,stringlength(11|11)"`
}

type DeleteUserDto struct {
	Id       int    `json:"id"`
	Password string `json:"password,omitempty"`
}

func marshalWithFormattedDate(v interface{}, dateFieldName string, dateValue time.Time) ([]byte, error) {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// Create a map to hold the field values
	m := make(map[string]interface{})
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		fieldValue := val.Field(i).Interface()
		fieldName := field.Tag.Get("json")
		if fieldName == "-" {
			continue
		}
		if fieldName == "" {
			fieldName = field.Name
		}
		m[fieldName] = fieldValue
	}

	// Set the formatted date
	m[dateFieldName] = dateValue.Format("2006-01-02")

	// Marshal the map to JSON
	return json.Marshal(m)
}

type FiltersDto struct {
	Name       string
	Surname    string
	Patronymic string
	TaskId     int
}

type FiltersMapDto struct {
	Filters map[string]interface{}
}
