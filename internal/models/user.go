package models

type User struct {
	Id             int64  `json:"id"  example:"13"`
	PassportNumber string `json:"passportNumber,omitempty" example:"1234 123456"`
	Name           string `json:"name,omitempty"  example:"Иван"`
	Surname        string `json:"surname,omitempty"  example:"Иванов"`
	Patronymic     string `json:"patronymic,omitempty"  example:"Иванович"`
	Address        string `json:"adress,omitempty" example:"г. Екатеринбург, ул. Берии, д. 7, кв. 3"`
}
