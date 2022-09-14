package signup

import (
	"time"
)

type Signup struct {
	ID          uint `gorm:"primaryKey"`
	Email       string 
	Username    string
	Password    string
	TimeCreated time.Time `gorm:"autoCreateTime"`
	TimeUpdated time.Time `gorm:"autoUpdateTime"`
}



type Login struct {
	Username string
	Hashpass string
	Email    string
}
