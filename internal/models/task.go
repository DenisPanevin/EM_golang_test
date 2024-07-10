package models

type Task struct {
	Id     int64  `json:"id"`
	Name   string `json:"name" valid:"required ,minstringlength(3)"`
	Status int    `json:"name" valid:"required , range(0|4)"`
}
