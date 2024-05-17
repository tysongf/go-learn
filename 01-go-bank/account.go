package main

import "time"

type transaction struct {
	Amount float64
	OccuredAt time.Time
}

type Account struct {
	Name string
	Number int64
	Transactions []transaction
}

func GetAccountBalance(account Account) float64 {
	balance := 0.0
	for _, trans := range account.Transactions {
		balance += trans.Amount
	}
	return balance
}