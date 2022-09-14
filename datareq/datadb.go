package datareq

import (
	"time"
)

type DBdata struct {
	ID         int `gorm:"primary_key"`
	Name       string
	Deposit    Deposit   `gorm:"embedded"`
	TimeCreate time.Time `gorm:"autoCreateTime"`
	TimeUpdate time.Time `gorm:"autoUpdateTime"`
}
