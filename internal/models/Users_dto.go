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
	Id             int
	PassportNumber string `json:"passportNumber" validate:"omitempty,passportNumber"`
	Name           string `json:"name" validate:"omitempty,min=5"`
	Surname        string `json:"surname" validate:"omitempty,min=3"`
	Patronymic     string `json:"patronymic" validate:"omitempty,min=3"`
	Address        string `json:"address" validate:"omitempty,min=28"`
}

type PassportNumberDto struct {
	PassportNumber string `json:"passportNumber" validate:"required,passportNumber"`
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
	UserId         int64             `json:"user_id"`
	Name           string            `json:"name"`
	Surname        string            `json:"surname"`
	Patronymic     string            `json:"patronymic"`
	Address        string            `json:"address"`
	PassportNumber string            `json:"passport_number"`
	Jobs           []ShowUserTaskDto `json:"jobs,omitempty"`
}

type ShowUserTaskDto struct {
	TaskId    int64  `json:"task_id"`
	TaskName  string `json:"task_name"`
	TotalWork Interval
}

type Interval struct {
	pgtype.Interval
}

type IntervalJSON struct {
	OnTask string `json:"time_on_task"`
}

func (i Interval) MarshalJSON() ([]byte, error) {
	micros := i.Microseconds

	minutes := micros / 6e+7
	hours := (minutes / 60) + int64(i.Days*24)
	minutes %= 60

	intervalJSON := IntervalJSON{
		OnTask: fmt.Sprintf("%vh : %vm", hours, minutes),
	}

	return json.Marshal(intervalJSON)
}
