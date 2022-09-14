package signup

type Signup_Req struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type Loginreq struct {
	Username string `json:"username"`
	Hashpass string `json:"password"`
	Email    string `json:"email"`
}