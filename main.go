package main

import (
	"bank/repository"
	"bank/service"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {

	db, err := sqlx.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/banking")
	if err != nil {
		log.Fatal(err)
	}

	custommerRepository := repository.NewCustomerRepositoryDB(db)

	custService := service.NewCustomerService(custommerRepository)

	customers, err := custService.GetCustomers()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(customers)
}
