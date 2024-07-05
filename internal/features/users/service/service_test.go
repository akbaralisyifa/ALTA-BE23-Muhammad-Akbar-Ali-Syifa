package service_test

import (
	"testing"
	"todos/internal/features/users"
	"todos/internal/features/users/service"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T){
	srv := service.NewUserServices()
	input := users.User{Username: "barre", Password: "1234", Email: "barre@gmail.com", Phone: "0812344322"}
	
	// sebelum menjalankan fungsi nya, maka jalan dulu method yang ada di dalam nya
	// method 
	// nam fungsi yang di panggil di dari query di service
	qry.On("Register", input).Return(nil).Once();
	
	err:= srv.Register(input);

	// kalau sukses balikan nilain nya
	assert.Nil(t, err);
}
