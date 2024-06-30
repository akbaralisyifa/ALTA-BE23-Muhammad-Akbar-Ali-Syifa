package users

import "todos/internal/models"

// struct untuk response data yang di tampilkan (di kembalikan)
type LoginResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"hp"`
	Token 	 string `json:"token"`
}

// buat fungsi untuk response yang di kembalikan dari login
func ToLoginResponse(input models.Users, token string) LoginResponse {
	return LoginResponse{
		ID		: input.ID,
		Username: input.Username,
		Email	: input.Email,
		Password: input.Password,
		Phone	: input.Phone,
		Token	: token,
	}
}

