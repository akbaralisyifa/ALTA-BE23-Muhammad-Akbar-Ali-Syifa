package users

// struct untuk penginputan yang perlu saja (yang perlu di isi)
// lalu ganti controller login nya menjadi truct ini bukan struct dari users
type LoginRequest struct{
	Email 		string	`json:"email"`
	Password 	string	`json:"password"`
}
