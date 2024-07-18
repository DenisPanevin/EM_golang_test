package models

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateUserDto struct {
	PassportNumber string `json:"passportNumber,omitempty" validate:"required,min=5"`
	Name           string `json:"name,omitempty" validate:"required,min=5"`
	Surname        string `json:"surname,omitempty" validate:"required,min=3"`
	Patronymic     string `json:"patronymic,omitempty" validate:"required,min=3"`
	Address        string `json:"address,omitempty" validate:"required ,min=28"`
}

type UpdateUserDto struct {
	Id             int    `swaggerignore:"true"`
	PassportNumber string `json:"passportNumber" validate:"omitempty,passportNumber" example:"1234 123456"`
	Name           string `json:"name" validate:"omitempty,min=5" example:"Иванов"`
	Surname        string `json:"surname" validate:"omitempty,min=3" example:"Иванов"`
	Patronymic     string `json:"patronymic" validate:"omitempty,min=3" example:"Иванович"`
	Address        string `json:"address" validate:"omitempty,min=28" example:"г. Екатеринбург, ул. Берии, д. 7, кв. 3"`
}

type PassportNumberDto struct {
	PassportNumber string `json:"passportNumber" validate:"required,passportNumber" example:"1234 123456"`
}

type UserFiltersDto struct {
	Id             int    `validate:"omitempty,checkFilterId"`
	Name           string `validate:"omitempty,min=3"`
	Surname        string `validate:"omitempty,min=3"`
	Patronymic     string `validate:"omitempty,min=3"`
	PassportNumber string `validate:"omitempty,passportNumber"`
	Address        string `validate:"omitempty,min=28"`
}

type PageFiltersDto struct {
	Limit  int `validate:"omitempty"`
	Offset int
	Page   int `validate:"omitempty"`
}

func (f *PageFiltersDto) CalculateOffset() {
	if f.Page >= 1 {
		f.Offset = (f.Page - 1) * f.Limit
	}

}

type ShowUserDto struct {
	UserId         int64             `json:"user_id" example:"1"`
	Name           string            `json:"name" example:"Иван"`
	Surname        string            `json:"surname" example:"Иванов"`
	Patronymic     string            `json:"patronymic" example:"Иванович"`
	Address        string            `json:"address" example:"г. Екатеринбург, ул. Берии, д. 7, кв. 3"`
	PassportNumber string            `json:"passport_number" example:"1234 123456"`
	TotalWorkTime  *TotalWorkTime    `json:"total_work,omitempty" swaggertype:"string" example:"1h 1m"`
	Jobs           []ShowUserTaskDto `json:"jobs,omitempty" `
}

type ShowUserTaskDto struct {
	TaskId    int64    `json:"task_id" example:"1"`
	TaskName  string   `json:"task_name" example:"Idle"`
	TotalWork Interval `json:"time_on_task" swaggertype:"string" example:"1h 1m"`
}

type Interval struct {
	pgtype.Interval `swaggerignore:"true"`
}

func (i Interval) MarshalJSON() ([]byte, error) {
	micros := i.Microseconds

	minutes := micros / 6e+7
	hours := (minutes / 60) + int64(i.Days*24)
	minutes %= 60

	return json.Marshal(fmt.Sprintf("%vh : %vm", hours, minutes))
}

type TotalWorkTime struct {
	pgtype.Interval
}

func (i TotalWorkTime) MarshalJSON() ([]byte, error) {

	micros := i.Microseconds

	minutes := micros / 6e+7
	hours := (minutes / 60) + int64(i.Days*24)
	minutes %= 60

	return json.Marshal(fmt.Sprintf("%vh : %vm", hours, minutes))
}
