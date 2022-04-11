package service

type CustomerResponse struct {
	CustomerID int    `json:"customerId"`
	Name       string `json:"name"`
	Status     int    `json:"status"`
}

type CustomerService interface {
	GetCustomers() ([]CustomerResponse, error)
	GetCustomerById(int) (*CustomerResponse, error)
}
