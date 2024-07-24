package models

type CreateTaskDto struct {
	Name string `json:"name" validate:"required,min=3"`
}
