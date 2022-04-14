package service

import (
	"bank/errs"
	"bank/logs"
	"bank/repository"
	"database/sql"
	"net/http"
)

type customerService struct {
	custRepo repository.CustomerRepository
}

func NewCustomerService(custRepo repository.CustomerRepository) customerService {
	return customerService{custRepo: custRepo}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := s.custRepo.GetAll()
	if err != nil {
		logs.Error(err)
	}

	custRes := []CustomerResponse{}

	for _, customer := range customers {
		custResponse := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		custRes = append(custRes, custResponse)
	}

	return custRes, nil
}

func (s customerService) GetCustomerById(id int) (*CustomerResponse, error) {
	customer, err := s.custRepo.GetById(id)
	if err != nil {

		if err == sql.ErrNoRows {
			logs.Error(err)
			return nil, errs.AppError{
				Code:    http.StatusNotFound,
				Message: "customer not found",
			}
		}
		logs.Error(err)
		return nil, errs.AppError{
			Code:    http.StatusInternalServerError,
			Message: "unexpected",
		}
	}

	custResponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}

	return &custResponse, nil
}
