package biz

import (
	"context"
	"github.com/muhadif/wa-bot-wrapper/internal/data"
	"github.com/muhadif/wa-bot-wrapper/pkg"
	"regexp"
	"strings"
)

type WAUseCase interface {
	ParseMessageAndDoCallback(ctx context.Context, payload string) (string, error)
}

type waUseCase struct {
	financialUseCase FinancialUseCase
}

func NewWAUseCase(financialUseCase FinancialUseCase) WAUseCase {
	return &waUseCase{financialUseCase: financialUseCase}
}

func (w waUseCase) ParseMessageAndDoCallback(ctx context.Context, payload string) (string, error) {
	uc, content := getExpression(payload)
	switch uc {
	case "create-financial-transaction":
		err := w.financialUseCase.CreateNewTransactionRequest(ctx, &data.CreateNewTransactionRequest{
			Category:    pkg.GetContentIfExist(content, 2),
			Amount:      pkg.GetContentIfExist(content, 3),
			Description: pkg.GetContentIfExist(content, 4),
		})
		if err != nil {
			return "", err
		}
		return "financial report added", err
	case "get-financial-transaction":
		dateQuery := pkg.GetContentIfExist(content, 2)
		dateFrom, dateTo := pkg.GetRangeDateByDateQuery(dateQuery)
		reply, err := w.financialUseCase.GetTransaction(ctx, &data.GetTransactionRequest{
			DateTo:   dateTo,
			DateFrom: dateFrom,
			Category: pkg.GetContentIfExist(content, 3),
		})
		if err != nil {
			return "", err
		}
		return reply.ToMessage(), err
	}
	return "command not found", nil
}

func getExpression(str string) (string, []string) {
	patternCreateFinancial := "f add"
	isMatch, _ := regexp.MatchString(patternCreateFinancial, str)
	if isMatch {
		content := strings.Split(str, " ")
		if len(content) > 3 {
			return "create-financial-transaction", content
		}
	}

	patternGetFinancial := "f show"
	isMatch, _ = regexp.MatchString(patternGetFinancial, str)
	if isMatch {
		content := strings.Split(str, " ")
		if len(content) > 2 {
			return "get-financial-transaction", content
		}
	}

	return "", nil
}
