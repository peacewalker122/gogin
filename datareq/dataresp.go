package datareq

type Dataresp struct{
	ID         int `json:"id"`
	Name       string `json:"name"`
	Deposit    Deposit   `json:"deposit"  gorm:"embedded"`
}