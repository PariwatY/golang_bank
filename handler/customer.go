package handler

import (
	"bank/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type customerHandler struct {
	cusService service.CustomerService
}

func NewCustomerHandler(cusService service.CustomerService) customerHandler {
	return customerHandler{cusService: cusService}
}

func (h customerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	customers, err := h.cusService.GetCustomers()
	if err != nil {
		handleError(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)

}

func (h customerHandler) GetCustomerById(w http.ResponseWriter, r *http.Request) {
	customerId, _ := strconv.Atoi(mux.Vars(r)["customerId"])
	customer, err := h.cusService.GetCustomerById(customerId)
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)

}
