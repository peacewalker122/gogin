package signup

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Signup(data Signup_Req) (Signup, error)
	Login(data Loginreq) error
}
type service struct {
	repo Repository
}

func Newservice(repo Repository) *service {
	return &service{repo}
}
func (s *service) Signup(data Signup_Req) (Signup, error) {
	test, _ := s.repo.signupvalidation(data.Username)
	test1, _ := s.repo.signupvalidation(data.Email)

	if test.Username != "" || test1.Email != "" {
		return Signup{}, errors.New("Username/Email already exist")
	}

	pwdhash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	user := Signup{
		Email:    data.Email,
		Username: data.Username,
		Password: string(pwdhash),
	}
	save, err := s.repo.Signup(user)
	return save, err

}
func (s *service) Login(data Loginreq) error {
	r, _ := s.repo.Login(Login(data))
	err := bcrypt.CompareHashAndPassword([]byte(r.Password), []byte(data.Hashpass))
	return err
}
