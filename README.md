# Instashop-e-commerce-api
A RESTful API for an e-commerce application. This API will handle basic CRUD operations for products and orders, and provide user management and authentication

## Features Implementation
* User Management (Register, Login)
* Product Management (Create,Delete,Update, Read)
* Order Management(Place Order, Get User Order,Cancel Order, Update Order Status)

## Application Dependencies
* Go 1.20 >=
* MySQL
* Gorm
* JWT
* Goose

## Usage
Make sure you have the MySQL service running on the machine as specified as the dependencies for this project. Change the configuration as it suit you in the .env

In the Shared/Database directory, run the Migration files to create database tables using this below command.Make sure you have goose binary installed and placed in the GOBIN path.Also there is a seed.sql data in the seed directory, run to populate against the table in the database

Goose Installation
```
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Migration
```
goose -dir shared/database/migrations -table _migrations mysql "root:P@ssw0r1d@tcp(localhost:3306)/ECommerceDb?charset=utf8mb4&parseTime=True&loc=Local" up
```

To run locally, cd from the root to the directory and run 

```
go run main.go serveapi --APP_ENV Development
```

To run on Docker
```
docker-compose -f docker-compose.dev.yml up -d
```

To Access the endpoint use this URL
```
http://localhost:8189/swagger/index.html
```


