package models

type Task struct {
	Id   int64  `json:"id" validate:"omitempty,gte=0"`
	Name string `json:"name" validate:"omitempty,min=3"`
}
