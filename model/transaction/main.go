package transaction

import "time"

//Transaction json
type Transaction struct {
	Title     string    `json:"title"`
	Amount    float32   `json:"amount"`
	Type      int       `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

//Transactions array of json
type Transactions []Transaction
