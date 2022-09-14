package signup

import "gorm.io/gorm"

type Repository interface {
	Signup(user Signup) (Signup, error)
	signupvalidation(target string)(Signup, error)
	Login(user Login) (Signup,error)
}
type repository struct {
	db *gorm.DB
}

func Newrepo(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository)Signup(user Signup) (Signup, error) {
	err := r.db.Create(&user).Error
	return user, err
}
func (r * repository) Login(user Login) (Signup,error) {
	var test Signup
	err := r.db.Where("username=? or email=?", user.Username,user.Email).Find(&test).Error
	return test,err
}
func (r *repository)signupvalidation(target string)(Signup, error) {
	var user Signup
	err := r.db.Where("username = ?", target).Find(&user).Error
	return user,err
}
