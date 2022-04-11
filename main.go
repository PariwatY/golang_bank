package main

import (
	"bank/handler"
	"bank/repository"
	"bank/service"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func main() {

	db, err := sqlx.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/banking")
	if err != nil {
		log.Fatal(err)
	}

	customerRepositoryDB := repository.NewCustomerRepositoryDB(db)
	customerRepositoryMock := repository.NewCustomerRepositoryMock()
	_ = customerRepositoryDB
	custService := service.NewCustomerService(customerRepositoryMock)
	customerHandler := handler.NewCustomerHandler(custService)

	router := mux.NewRouter()

	router.HandleFunc("/customers", customerHandler.GetCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerId:[0-9]+}", customerHandler.GetCustomerById).Methods(http.MethodGet)

	http.ListenAndServe(":8080", router)
	// customers, err := custService.GetCustomers()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(customers)

	customer, err := custService.GetCustomerById(2001)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(customer)
}
