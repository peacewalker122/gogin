package datareq

type Userreq struct {
	Name    string  `json:"name" binding:"required,alpha"`
	Deposit Deposit `json:"deposit"`
}
type Deposit struct {
	Amount   int    `json:"amount" binding:"required,numeric"`
	Currency string `json:"currency" binding:"required,alpha"`
}

// for htnl box
