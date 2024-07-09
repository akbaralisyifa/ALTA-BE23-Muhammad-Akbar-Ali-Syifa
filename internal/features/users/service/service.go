package service

import (
	"todos/internal/features/users"
	"todos/internal/utils"
)

type UserServices struct {
	qry users.Query
	jwt utils.JwtUtilityInterface
	vldt utils.ValidateUtilityInterface
	pw utils.GeneretePasswordInterface
}

func NewUserServices(m users.Query, v utils.ValidateUtilityInterface, j utils.JwtUtilityInterface, p utils.GeneretePasswordInterface) users.Services {
	return &UserServices{
		qry: m,
		vldt: v,
		jwt: j,
		pw : p,
	}
}

func (us *UserServices) Register(newData users.User)(error){

	err := us.vldt.RegisterValidator(newData.Username, newData.Email, newData.Password)

	if err != nil {
		return err;
	}

	processPw, err := us.pw.GeneretePassword(newData.Password)

	if err != nil {
		return err;
	}

	newData.Password = string(processPw);

	// mengecek ketika input nya sudah di isi namun tidak tersimpan ke server
	err = us.qry.Register(newData);
	if err != nil {
		return err;
	}

	return nil;
};

func (us *UserServices) Login(email string, password string)(users.User, string, error){

	err := us.vldt.LoginValidator(email, password)

	if err != nil {
		return users.User{}, "", err;
	}

	// data yang di input kan
	result, err := us.qry.Login(email);

	if err != nil {
		return users.User{}, "", err
	}

	// cek password
	err = utils.CheckPassword([]byte(password), []byte(result.Password));

	if err != nil {
		return users.User{}, "", err
	}

	// generet token
	token, err := us.jwt.GenereteToken(result.ID)

	if err != nil {
		return users.User{}, "", err
	}

	return result, token, nil;
}