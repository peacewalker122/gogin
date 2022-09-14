package signup

type Response struct{
	ID uint `json:"id"`
	Email string `json:"email"`
	Username string `json:"username"`
}