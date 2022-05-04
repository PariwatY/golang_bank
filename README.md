# About project
- Create Project using Hexagonal Architecture for go project 

# List Task in this project
1. Create Database adapter to connect with Mysql
2. Create Get API with Mux
  - Get customers
  - Get customer by id
3. Configuration with Viper Ex: database config etc.


# Lib (only important lib that i thought haha)
1. github.com/go-sql-driver/mysql
 - Using for connect with MySQL
2. github.com/gorilla/mux
 - Using for creating API
3. github.com/jmoiron/sqlx
 - Using for reducing a syntax when query data

 
# Curl for testing API
1. Create account
```
curl localhost:8081/customers/2003/accounts -i -X POST -H  "content-type:application/json" -d '{"account_type":"saving","amount":29999.83}'
```
2. Get account by customer id
```
curl localhost:8081/customers/2001/accounts -i
```
3. Test amount < 5000
```
curl localhost:8081/customers/2004/accounts -i -X POST -H  "content-type:application/json" -d '{"account_type":"saving","amount":200}'
```
4. Test accoun type  is not saving and checking
```
curl localhost:8081/customers/2004/accounts -i -X POST -H  "content-type:application/json" -d '{"account_type":"ss","amount":5000}'
```
