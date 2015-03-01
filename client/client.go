package client

import (
	"errors"

	"jam/vendors/client/zygote"
)

type Register struct {
	Name, Password, Email string
}

func (r Register) Now() error {
	var e error = nil

	if zygote.ByName(r.Name) != nil {
		e = errors.New("client: Unable to register user with the name '" + r.Name + "'")
	}

	if zygote.ByEmail(r.Email) != nil {
		e = errors.New("client: Unable to register user with the email '" + r.Email + "'")
	}

	if e == nil {
		zygote.Data {
			Name: r.Name,
			Password: r.Password,
			Email: r.Email,	
		}.Insert()
	}

	return e
}
