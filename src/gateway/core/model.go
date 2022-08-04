package core

import (
	"fmt"
	"strings"
)

var (
	ErrUsernameLeng = fmt.Errorf("username contains at least 6 and maximal 15 character")
	ErrPassLeng     = fmt.Errorf("password contains at least 8 and maximal 32 character")
	ErrInvalidChar  = fmt.Errorf("invalid character")
)

type IRequest interface {
	Valid() error
}

func ValidateRequest(req IRequest) error {
	return req.Valid()
}

type Authentication struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *Authentication) Valid() error {
	r.Username = strings.ToLower(Trim(r.Username))
	r.Password = Trim(r.Password)

	if !IsAlphaNumeric(r.Username) {
		return ErrInvalidChar
	}

	for _, v := range r.Password {
		// Invalid character \ ' "
		if v == 34 || v == 92 || v == 39 {
			return ErrInvalidChar
		}
	}

	if lengRule(6, 15, r.Username) {
		return ErrUsernameLeng
	}

	if lengRule(8, 32, r.Password) {
		return ErrPassLeng
	}

	return nil
}

type Name struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func (r Name) Valid() error {
	if !IsAlphaNumeric(r.Firstname) {
		return ErrInvalidChar
	}

	if !IsAlphaNumeric(r.Lastname) {
		return ErrInvalidChar
	}

	if lengRule(3, 10, r.Firstname) {
		return ErrUsernameLeng
	}

	if lengRule(3, 15, r.Lastname) {
		return ErrUsernameLeng
	}

	return nil
}

type RegisterUser struct {
	Authentication
	Email string `json:"email"`
	Name
}

func (r *RegisterUser) Valid() error {
	if err := r.Authentication.Valid(); err != nil {
		return err
	}

	if err := r.Name.Valid(); err != nil {
		return err
	}

	return nil
}
