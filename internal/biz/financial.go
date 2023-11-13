package biz

import (
	"context"
	"github.com/muhadif/wa-bot-wrapper/internal/data"
	"github.com/muhadif/wa-bot-wrapper/internal/repository"
)

type FinancialUseCase interface {
	CreateNewTransactionRequest(ctx context.Context, payload *data.CreateNewTransactionRequest) error
	GetTransaction(ctx context.Context, req *data.GetTransactionRequest) (*data.GetTransactionResponse, error)
}

type financialUseCase struct {
	financialRepo repository.FinancialRepository
}

func NewFinancialUseCase(financialRepo repository.FinancialRepository) FinancialUseCase {
	return &financialUseCase{financialRepo: financialRepo}
}

func (f financialUseCase) CreateNewTransactionRequest(ctx context.Context, payload *data.CreateNewTransactionRequest) error {
	return f.financialRepo.CreateNewTransaction(ctx, payload)
}

func (f financialUseCase) GetTransaction(ctx context.Context, req *data.GetTransactionRequest) (*data.GetTransactionResponse, error) {
	return f.financialRepo.GetTransaction(ctx, req)
}
