package models

type CreateTaskDto struct {
	Name   string `json:"name" valid:"required ,minstringlength(3)"`
	Status int    `json:"status" valid:"required , range(1|4)"`
}
