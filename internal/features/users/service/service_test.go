package service_test

import (
	"errors"
	"testing"
	"todos/internal/features/users"
	"todos/internal/features/users/service"
	"todos/mocks"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func TestRegister(t *testing.T){
	qry := mocks.NewQuery(t);
	vldt:= mocks.NewValidateUtilityInterface(t)
	jwt := mocks.NewJwtUtilityInterface(t)
	pw := mocks.NewGeneretePasswordInterface(t)
	srv := service.NewUserServices(qry, vldt, jwt, pw);

	input := users.User{Username: "akbaralisyifa", Password: "akbar12345", Email: "akbaralisyifa@gmail.com", Phone: "087655343"}

	t.Run("Success Register", func(t *testing.T) {
		inputQry := users.User{Username: "akbaralisyifa", Password: "somepassword", Email: "akbaralisyifa@gmail.com", Phone: "087655343"}

		// Expecting 4 arguments for RegisterValidator
		vldt.On("RegisterValidator", input.Username, input.Email, input.Password).Return(nil).Once()
		pw.On("GeneretePassword", input.Password).Return([]byte("somepassword"), nil).Once()
		qry.On("Register", inputQry).Return(nil).Once()

		err := srv.Register(input)

		vldt.AssertExpectations(t)
		jwt.AssertExpectations(t)
		pw.AssertExpectations(t)
		qry.AssertExpectations(t)

		assert.Nil(t, err)
	});
	t.Run("Error Hash Password", func(t *testing.T) {
		originalPassword := input.Password
		vldt.On("RegisterValidator", input.Username, input.Email, input.Password, input.Phone).Return(nil).Once()
		pw.On("GeneretePassword", input.Password).Return(nil, bcrypt.ErrPasswordTooLong).Once()

		err := srv.Register(input)

		vldt.AssertExpectations(t)
		pw.AssertExpectations(t)

		assert.Error(t, err)
		assert.ErrorContains(t, err, bcrypt.ErrPasswordTooLong.Error())
		assert.Equal(t, originalPassword, input.Password) 

	})

	t.Run("Error From Query", func(t *testing.T) {
		inputQry := users.User{Username: "akbaralisyifa", Password: "somepassword", Email: "akbaralisyifa@gmail.com", Phone: "087655343"}
		vldt.On("RegisterValidator", input.Username, input.Email, input.Password, input.Phone).Return(nil).Once()
		pw.On("GeneretePassword", input.Password).Return([]byte("goodpassword"), nil).Once()
		qry.On("Register", inputQry).Return(gorm.ErrInvalidData).Once()

		err := srv.Register(input)

		vldt.AssertExpectations(t)
		pw.AssertExpectations(t)
		qry.AssertExpectations(t)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "terjadi kesalahan pada server saat mengolah data")

	})

	t.Run("Nil Value", func(t *testing.T) {
		vldt.On("RegisterValidator", input.Username, input.Email, input.Password).Return(errors.New("validasi tidak sesuai")).Once()

		err := srv.Register(input)

		vldt.AssertExpectations(t)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "validasi tidak sesuai")
	})
};


func TestLogin(t *testing.T){
	qry := mocks.NewQuery(t);
	vldt:= mocks.NewValidateUtilityInterface(t)
	jwt := mocks.NewJwtUtilityInterface(t)
	pw := mocks.NewGeneretePasswordInterface(t)
	srv := service.NewUserServices(qry, vldt, jwt, pw);

	input := users.User{Password: "akbar12345", Email: "akbaralisyifa@gmail.com"};

	t.Run("Success Login", func(t *testing.T) {
		inputQry := users.User{Password: "akbar12345", Email: "akbaralisyifa@gmail.com"}

		vldt.On("EmailPasswordValidator", input.Email, input.Password).Return(nil).Once()
		qry.On("Login", input.Email).Return(inputQry, nil).Once()
		pw.On("CheckPassword", []byte(inputQry.Password), []byte(inputQry.Password)).Return(nil).Once()
		jwt.On("GenerateJWT", input.ID, input.Email).Return("someToken", nil).Once()

		user, token, err := srv.Login(inputQry.Email, inputQry.Password)

		vldt.AssertExpectations(t)
		qry.AssertExpectations(t)
		pw.AssertExpectations(t)
		jwt.AssertExpectations(t)

		assert.Nil(t, err)
		assert.Equal(t, "someToken", token)
		assert.Equal(t, inputQry, user)
	})

	t.Run("Failed email/password Login", func(t *testing.T) {
		inputQry := users.User{Password: "", Email: ""}

		vldt.On("EmailPasswordValidator", inputQry.Email, inputQry.Password).Return(errors.New("validasi gagal")).Once()

		_, _, err := srv.Login(inputQry.Email, inputQry.Password)

		vldt.AssertExpectations(t)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "validasi gagal")
	})

	t.Run("Error on Query", func(t *testing.T) {
		inputQry := users.User{Password: "anggi1234", Email: "anggi@eko.com"}

		vldt.On("EmailPasswordValidator", inputQry.Email, inputQry.Password).Return(nil).Once()
		qry.On("Login", input.Email).Return(inputQry, gorm.ErrInvalidData).Once()

		_, _, err := srv.Login(inputQry.Email, inputQry.Password)

		vldt.AssertExpectations(t)
		qry.AssertExpectations(t)

		assert.Error(t, err)
		assert.ErrorContains(t, err, string(gorm.ErrInvalidData.Error()))
	})

	t.Run("Error on Password", func(t *testing.T) {
		inputQry := users.User{Password: "akbar12345", Email: "akbaralisyifa@gmail.com"}

		vldt.On("EmailPasswordValidator", inputQry.Email, inputQry.Password).Return(nil).Once()
		qry.On("Login", input.Email).Return(inputQry, nil).Once()
		pw.On("CheckPassword", []byte(inputQry.Password), []byte(inputQry.Password)).Return(bcrypt.ErrMismatchedHashAndPassword).Once()

		_, _, err := srv.Login(inputQry.Email, inputQry.Password)

		vldt.AssertExpectations(t)
		qry.AssertExpectations(t)
		pw.AssertExpectations(t)

		assert.Error(t, err)
		assert.ErrorContains(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
	})

	t.Run("Error on JWT Generator", func(t *testing.T) {
		inputQry := users.User{Password: "akbar12345", Email: "akbaralisyifa@gmail.com"}

		vldt.On("EmailPasswordValidator", inputQry.Email, inputQry.Password).Return(nil).Once()
		qry.On("Login", input.Email).Return(inputQry, nil).Once()
		pw.On("CheckPassword", []byte(inputQry.Password), []byte(inputQry.Password)).Return(nil).Once()
		jwt.On("GenerateJWT", input.ID).Return("", errors.New("Tidak dapat mendapatkan token")).Once()

		_, _, err := srv.Login(inputQry.Email, inputQry.Password)

		vldt.AssertExpectations(t)
		qry.AssertExpectations(t)
		pw.AssertExpectations(t)
		jwt.AssertExpectations(t)

		assert.Error(t, err)
		assert.ErrorContains(t, err, "Tidak dapat mendapatkan token")
	})


}
