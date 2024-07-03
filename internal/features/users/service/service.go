package service

import (
	"todos/internal/features/users"
	"todos/internal/utils"
)

type UserServices struct {
	qry users.Query
}

func NewUserServices(m users.Query) users.Services {
	return &UserServices{
		qry: m,
	}
}

func (us *UserServices) Register(newData users.User)(error){
	processPw, err := utils.GeneretePassword(newData.Password);

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
	token, err := utils.GenereteToken(result.ID);

	if err != nil {
		return users.User{}, "", err
	}

	return result, token, nil;
}