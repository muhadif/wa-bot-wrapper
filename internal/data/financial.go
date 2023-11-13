package data

import (
	"fmt"
	"time"
)

type CreateNewTransactionRequest struct {
	Category    string
	Amount      string
	Description string
}

type GetTransactionRequest struct {
	DateTo   *time.Time
	DateFrom *time.Time
	Category string
}

type GetTransactionResponse struct {
	DateTo      *time.Time
	DateFrom    *time.Time
	Transaction []*TransactionPerCategory
}

func (r GetTransactionResponse) ToMessage() string {
	content := fmt.Sprintf("Date from : %s \n", r.DateFrom.Format(time.DateOnly))
	content += fmt.Sprintf("Date to : %s \n", r.DateTo.Format(time.DateOnly))
	content += "Category  |  Amount \n"
	for _, transaction := range r.Transaction {
		content += fmt.Sprintf("%s | Rp %d \n", transaction.Category, transaction.Amount)
	}
	return content
}

type TransactionPerCategory struct {
	Amount   int
	Category string
}
