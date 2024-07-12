package models

import (
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
	PassportNumber *string `json:"passportNumber,omitempty" validate:"min=5"`
	Name           *string `json:"name,omitempty" validate:"min=5"`
	Surname        *string `json:"surname,omitempty" validate:"min=3"`
	Patronymic     *string `json:"patronymic,omitempty" validate:"min=3"`
	Address        *string `json:"address,omitempty" validate:" min=28"`
}

type PassportNumberDto struct {
	PassportNumber string `json:"passportNumber" validate:"required,passportNumber"`
}

type DeleteUserDto struct {
	Id       int    `json:"id"`
	Password string `json:"password,omitempty"`
}

// Create a map to hold the field values

type FiltersDto struct {
	Name       string
	Surname    string
	Patronymic string
	TaskId     int
	Limit      int
	Page       int
}

type ShowUserDto struct {
	UserId int64 `json:"user_id"`

	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`

	Address        string            `json:"address"`
	PassportNumber string            `json:"passport_number"`
	TotalWork      []ShowUserTaskDto `json:"total_work"`
}

type ShowUserTaskDto struct {
	TaskId    int64  `json:"task_id"`
	TaskName  string `json:"-"`
	TotalWork Interval
}

type Interval struct {
	pgtype.Interval
}

type IntervalJSON struct {
	Hours   int64 `json:"hours"`
	Minutes int64 `json:"minutes"`
}

func (i Interval) MarshalJSON() ([]byte, error) {
	var totalMinutes int64

	// Convert interval to total minutes

	hours := totalMinutes / 60
	minutes := totalMinutes % 60

	intervalJSON := IntervalJSON{
		Hours:   hours,
		Minutes: minutes,
	}

	return json.Marshal(intervalJSON)
}
