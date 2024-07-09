package utils

import "golang.org/x/crypto/bcrypt";


type GeneretePasswordInterface interface{
	GeneretePassword(currentPw string) ([]byte, error) 
};

type passwordUtility struct{};

func NewGenertePassword() GeneretePasswordInterface {
	return &passwordUtility{};
}

// fungsi ketika user register (generete pw nya)
func (pw *passwordUtility) GeneretePassword(currentPw string) ([]byte, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(currentPw), bcrypt.DefaultCost);

	if err != nil {
		return nil, err
	}

	return result, nil;
}

func CheckPassword(inputPw []byte, currentPw []byte)(error) {
	return	bcrypt.CompareHashAndPassword(currentPw, inputPw)
}