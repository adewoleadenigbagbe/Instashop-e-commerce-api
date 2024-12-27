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

## Usage
Make sure you have the MySQL service running on the machine as specified as the dependencies for this project. Change the configuration as it suit you in the .env

To run locally, cd from the root to the directory and run 

```
go run main.go serveapi
```

To Access the endpoint use this URL
```
http://localhost:8189/swagger/index.html
```


