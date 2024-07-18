package pkg

import "errors"

var (
	ApiWrongInput  = errors.New("Wrong input")
	CantCreateUser = errors.New("Cant create user")
	AlreadyExists  = errors.New("Already Exists")
	NotFound       = errors.New("Not Found")
)
