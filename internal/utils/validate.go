package utils

import (
	"todos/internal/features/users"

	"github.com/go-playground/validator/v10"
)

type ValidateUtilityInterface interface {
	RegisterValidator(username string, email string, password string) error
	LoginValidator(email string, password string) error
}

type validateUtility struct {
	vldt validator.Validate
}

func NewValidatorUtility(v validator.Validate) ValidateUtilityInterface {
	return &validateUtility{
		vldt: v,
	}
}

func(vu *validateUtility) RegisterValidator(username string, email string, password string) error {
	err := vu.vldt.Struct(&users.RegisterValidate{Username: username, Email: email, Password: password})

	if err != nil {
		return err;
	}
	return nil;
}

func(vu *validateUtility) LoginValidator(email string, password string) error {
	err := vu.vldt.Struct(&users.LoginValidate{Email: email, Password: password})

	if err != nil {
		return err;
	}
	return nil;
}