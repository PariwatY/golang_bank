package service

import (
	"bank/repository"
	"log"
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
		log.Println(err)
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
	return nil, nil
}
