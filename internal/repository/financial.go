package repository

import (
	"context"
	"github.com/google/martian/log"
	"github.com/muhadif/wa-bot-wrapper/internal/data"
	"gorm.io/gorm"
	"time"
)

type FinancialRepository interface {
	CreateNewTransaction(ctx context.Context, req *data.CreateNewTransactionRequest) error
	GetTransaction(ctx context.Context, req *data.GetTransactionRequest) (*data.GetTransactionResponse, error)
}

type financialRepository struct {
	db *gorm.DB
}

func NewFinancialRepository(db *gorm.DB) FinancialRepository {
	return &financialRepository{db: db}
}

func (f financialRepository) CreateNewTransaction(ctx context.Context, req *data.CreateNewTransactionRequest) error {
	if err := f.db.Table("financial_transaction").Create(req); err != nil {
		log.Errorf("CreateNewTransaction %v", err)
	}

	return nil
}

func (f financialRepository) GetTransaction(ctx context.Context, req *data.GetTransactionRequest) (*data.GetTransactionResponse, error) {
	query := f.db.Debug().Table("financial_transaction").
		Where("DATE(transaction_date) >= ?", req.DateFrom.Format(time.DateOnly)).
		Where("DATE(transaction_date) <= ?", req.DateTo.Format(time.DateOnly))

	if req.Category != "" {
		query = query.Where("category = ?", req.Category)
	}

	query = query.Select("category, SUM(amount) as amount").Group("category")

	var transactions []*data.TransactionPerCategory
	if err := query.Find(&transactions).Error; err != nil {
		log.Errorf("GetTransaction %v", err)
	}
	return &data.GetTransactionResponse{
		DateTo:      req.DateTo,
		DateFrom:    req.DateFrom,
		Transaction: transactions,
	}, nil
}
