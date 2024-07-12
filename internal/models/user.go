package models

type User struct {
	Id             int64  `json:"id"`
	PassportNumber string `json:"passportNumber,omitempty"`
	Name           string `json:"name,omitempty"`
	Surname        string `json:"surname,omitempty"`
	Patronymic     string `json:"patronymic,omitempty"`
	Address        string `json:"adress,omitempty"`
}
