package main

import (
	"bank/handler"
	"bank/logs"
	"bank/repository"
	"bank/service"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func main() {
	// Initial time zone to be Asia/Bangkok
	initTimeZone()
	// Initial Config data for this project such as Database Name, Database Port, App Port etc.
	initConfig()
	// Set database config
	db := initDB()

	customerRepositoryDB := repository.NewCustomerRepositoryDB(db)
	custService := service.NewCustomerService(customerRepositoryDB)
	customerHandler := handler.NewCustomerHandler(custService)

	accountRepositoryDB := repository.NewAccountRepositoryDB(db)
	accService := service.NewAccountService(accountRepositoryDB)
	accountHandler := handler.NewAccountHandler(accService)
	_ = accountHandler

	router := mux.NewRouter()

	// Customer Handler
	router.HandleFunc("/customers", customerHandler.GetCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerId:[0-9]+}", customerHandler.GetCustomerById).Methods(http.MethodGet)

	// Account Handler
	router.HandleFunc("/customers/{customerId:[0-9]+}/accounts", accountHandler.GetAccountByCustId).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerId:[0-9]+}/accounts", accountHandler.NewAccount).Methods(http.MethodPost)

	logs.Info("Bank Service started at port:" + viper.GetString("app.port"))
	http.ListenAndServe(fmt.Sprintf(":%v", viper.GetString("app.port")), router)

}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}

func initDB() *sqlx.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.database"))
	db, err := sqlx.Open(viper.GetString("db.driver"), dsn)
	if err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db

}
