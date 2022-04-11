package main

import (
	"bank/repository"
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

	// customers, err := custommerRepository.GetAll()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Print(customers)

	customer, err := custommerRepository.GetById(2002)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(customer)

}
