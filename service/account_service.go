package service

import (
	"bank/errs"
	"bank/logs"
	"bank/repository"
	"strings"
	"time"
)

type accountService struct {
	accRepo repository.AccountRepository
}

func NewAccountService(accRepo repository.AccountRepository) accountService {
	return accountService{accRepo: accRepo}
}

func (s accountService) NewAccount(customerID int, req NewAccountRequest) (*AccountResponse, error) {
	//TODO Validate
	if req.Amount < 5000 {
		return nil, errs.NewValidationError("amount must be greater than 5000")
	}

	if strings.ToLower(req.AccountType) != "saving" && strings.ToLower(req.AccountType) != "checking" {
		return nil, errs.NewValidationError("account type must be 'saving' or 'checking'")
	}

	account := repository.Account{
		CustomerID:  customerID,
		OpeningDate: time.Now().Format("2006-1-2 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      1,
	}
	accCreaRes, err := s.accRepo.Create(account)
	if err != nil {
		return nil, errs.NewUnExpectedError()
	}

	responses := AccountResponse{
		AccountID:   accCreaRes.AccountID,
		OpeningDate: accCreaRes.OpeningDate,
		AccountType: accCreaRes.AccountType,
		Amount:      accCreaRes.Amount,
		Status:      accCreaRes.Status,
	}

	return &responses, nil
}

func (s accountService) GetAccount(customerId int) ([]AccountResponse, error) {
	accounts, err := s.accRepo.GetAll(customerId)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewNotFoundError("Account not found")
	}

	responses := []AccountResponse{}
	for _, account := range accounts {
		responses = append(responses, AccountResponse{
			AccountID:   account.AccountID,
			OpeningDate: account.OpeningDate,
			AccountType: account.AccountType,
			Amount:      account.Amount,
			Status:      account.Status,
		})
	}

	return responses, nil
}
