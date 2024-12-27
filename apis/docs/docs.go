// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "get the status of server.",
                "consumes": [
                    "*/*"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "root"
                ],
                "summary": "Show the status of server.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/auth/register": {
            "post": {
                "description": "Register new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Register new user",
                "parameters": [
                    {
                        "description": "CreateUserRequest",
                        "name": "CreateUserRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services.CreateUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/auth/signin": {
            "post": {
                "description": "SignIn user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "SignIn user",
                "parameters": [
                    {
                        "description": "SignInRequest",
                        "name": "SignInRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.SignInRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services.SignInResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/orders": {
            "get": {
                "description": "Get all user order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order"
                ],
                "summary": "Get all user order",
                "parameters": [
                    {
                        "description": "GetOrderRequest",
                        "name": "GetOrderRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.GetOrderRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services.GetOrderResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a a new customer order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order"
                ],
                "summary": "Create a order",
                "parameters": [
                    {
                        "description": "CreateOrderRequest",
                        "name": "CreateOrderRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.CreateOrderRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/services.CreateOrderResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/orders/{id}": {
            "put": {
                "description": "Update user order status",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order"
                ],
                "summary": "Update user order status",
                "parameters": [
                    {
                        "description": "UpdateStatusRequest",
                        "name": "UpdateStatusRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.UpdateStatusRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/services.UpdateStatusResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/orders/{id}/cancel": {
            "put": {
                "description": "Cancel user order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order"
                ],
                "summary": "Cancel user order",
                "parameters": [
                    {
                        "description": "CancelOrderRequest",
                        "name": "CancelOrderRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.CancelOrderRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/services.CancelOrderResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/products": {
            "get": {
                "description": "Get a products",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "Get a products",
                "parameters": [
                    {
                        "description": "GetProductsRequest",
                        "name": "GetProductsRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.GetProductsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services.Result"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "Create a new product",
                "parameters": [
                    {
                        "description": "CreateProductRequest",
                        "name": "CreateProductRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.CreateProductRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/services.CreateProductResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/products/{id}": {
            "get": {
                "description": "Get a product by the ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "Get a product by the ID",
                "parameters": [
                    {
                        "description": "GetProductRequestById",
                        "name": "GetProductRequestById",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.GetProductRequestById"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services.GetProductByIdResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Update product with product info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "Update product with product info",
                "parameters": [
                    {
                        "description": "UpdateProductRequest",
                        "name": "UpdateProductRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.UpdateProductRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/services.UpdateProductResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a product by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "product"
                ],
                "summary": "Delete a product by id",
                "parameters": [
                    {
                        "description": "DeleteProductRequest",
                        "name": "DeleteProductRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.DeleteProductRequest"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/services.Result"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/user/add-role": {
            "post": {
                "description": "Create a new role",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "userole"
                ],
                "summary": "Create a new role",
                "parameters": [
                    {
                        "description": "CreateRoleRequest",
                        "name": "CreateRoleRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/services.CreateRoleRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/services.CreateRoleResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "enums.CategoryType": {
            "type": "integer",
            "enum": [
                1,
                2,
                3,
                4,
                5,
                6,
                7,
                8,
                9,
                10
            ],
            "x-enum-varnames": [
                "Electronics",
                "Fashion",
                "Beauty",
                "Sports",
                "Automotive",
                "Health",
                "Beverages",
                "Books",
                "Toys",
                "Home"
            ]
        },
        "enums.OrderStatus": {
            "type": "integer",
            "enum": [
                1,
                2,
                3,
                4,
                5
            ],
            "x-enum-varnames": [
                "Pending",
                "Confirmed",
                "Shipped",
                "Delivered",
                "Cancelled"
            ]
        },
        "enums.ProductStatus": {
            "type": "integer",
            "enum": [
                1,
                2,
                3
            ],
            "x-enum-varnames": [
                "Active",
                "InActive",
                "Discontinued"
            ]
        },
        "enums.Role": {
            "type": "integer",
            "enum": [
                1,
                2
            ],
            "x-enum-varnames": [
                "EndUser",
                "Admin"
            ]
        },
        "services.CancelOrderRequest": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "services.CancelOrderResponse": {
            "type": "object",
            "properties": {
                "orderId": {
                    "type": "string"
                }
            }
        },
        "services.CreateOrderRequest": {
            "type": "object",
            "required": [
                "orderItems",
                "userId"
            ],
            "properties": {
                "orderItems": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/services.OrderItemModel"
                    }
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "services.CreateOrderResponse": {
            "type": "object",
            "properties": {
                "orderId": {
                    "type": "string"
                }
            }
        },
        "services.CreateProductRequest": {
            "type": "object",
            "required": [
                "categoryId",
                "description",
                "imageUrl",
                "name",
                "price",
                "stock"
            ],
            "properties": {
                "categoryId": {
                    "type": "string"
                },
                "description": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 3
                },
                "imageUrl": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 3
                },
                "price": {
                    "type": "number"
                },
                "status": {
                    "$ref": "#/definitions/enums.ProductStatus"
                },
                "stock": {
                    "type": "integer",
                    "minimum": 0
                }
            }
        },
        "services.CreateProductResponse": {
            "type": "object",
            "properties": {
                "productId": {
                    "type": "string"
                }
            }
        },
        "services.CreateRoleRequest": {
            "type": "object",
            "required": [
                "description",
                "role"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "role": {
                    "$ref": "#/definitions/enums.Role"
                }
            }
        },
        "services.CreateRoleResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "services.CreateUserRequest": {
            "type": "object",
            "required": [
                "email",
                "firstName",
                "lastName",
                "password",
                "phoneNumber",
                "roleId"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phoneNumber": {
                    "type": "string"
                },
                "roleId": {
                    "type": "string"
                }
            }
        },
        "services.CreateUserResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "services.DeleteProductRequest": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "services.GetOrderItemModel": {
            "type": "object",
            "properties": {
                "categoryName": {
                    "type": "string"
                },
                "orderItemId": {
                    "type": "string"
                },
                "productDescription": {
                    "type": "string"
                },
                "productName": {
                    "type": "string"
                },
                "productPrice": {
                    "type": "number"
                },
                "quantity": {
                    "type": "integer"
                },
                "subTotal": {
                    "type": "number"
                }
            }
        },
        "services.GetOrderRequest": {
            "type": "object",
            "required": [
                "userId"
            ],
            "properties": {
                "order": {
                    "type": "string"
                },
                "page": {
                    "type": "integer"
                },
                "pageLength": {
                    "type": "integer"
                },
                "sortBy": {
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "services.GetOrderResponse": {
            "type": "object",
            "properties": {
                "orders": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/services.OrderModel"
                    }
                }
            }
        },
        "services.GetProductByIdResponse": {
            "type": "object",
            "properties": {
                "categoryDescription": {
                    "type": "string"
                },
                "categoryName": {
                    "type": "string"
                },
                "categoryType": {
                    "$ref": "#/definitions/enums.CategoryType"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "imageUrl": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/enums.ProductStatus"
                },
                "stock": {
                    "type": "integer"
                }
            }
        },
        "services.GetProductRequestById": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "services.GetProductResponse": {
            "type": "object",
            "properties": {
                "categoryDescription": {
                    "type": "string"
                },
                "categoryName": {
                    "type": "string"
                },
                "categoryType": {
                    "$ref": "#/definitions/enums.CategoryType"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "imageUrl": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/enums.ProductStatus"
                },
                "stock": {
                    "type": "integer"
                }
            }
        },
        "services.GetProductsRequest": {
            "type": "object",
            "properties": {
                "order": {
                    "type": "string"
                },
                "page": {
                    "type": "integer"
                },
                "pageLength": {
                    "type": "integer"
                },
                "sortBy": {
                    "type": "string"
                }
            }
        },
        "services.OrderItemModel": {
            "type": "object",
            "required": [
                "productId",
                "quantity"
            ],
            "properties": {
                "productId": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "services.OrderModel": {
            "type": "object",
            "properties": {
                "orderId": {
                    "type": "string"
                },
                "orderitems": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/services.GetOrderItemModel"
                    }
                },
                "taxAmount": {
                    "type": "number"
                },
                "totalPrice": {
                    "type": "number"
                }
            }
        },
        "services.Result": {
            "type": "object",
            "properties": {
                "page": {
                    "type": "integer"
                },
                "perPage": {
                    "type": "integer"
                },
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/services.GetProductResponse"
                    }
                },
                "requestedPageLength": {
                    "type": "integer"
                },
                "totalPage": {
                    "type": "integer"
                },
                "totalResults": {
                    "type": "integer"
                }
            }
        },
        "services.SignInRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "services.SignInResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                }
            }
        },
        "services.UpdateProductRequest": {
            "type": "object",
            "required": [
                "categoryId",
                "description",
                "imageUrl",
                "name",
                "price",
                "stock"
            ],
            "properties": {
                "categoryId": {
                    "type": "string"
                },
                "description": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 3
                },
                "id": {
                    "type": "string"
                },
                "imageUrl": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "maxLength": 100,
                    "minLength": 3
                },
                "price": {
                    "type": "number"
                },
                "status": {
                    "$ref": "#/definitions/enums.ProductStatus"
                },
                "stock": {
                    "type": "integer",
                    "minimum": 0
                }
            }
        },
        "services.UpdateProductResponse": {
            "type": "object",
            "properties": {
                "productId": {
                    "type": "string"
                }
            }
        },
        "services.UpdateStatusRequest": {
            "type": "object",
            "required": [
                "id",
                "status"
            ],
            "properties": {
                "id": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/enums.OrderStatus"
                }
            }
        },
        "services.UpdateStatusResponse": {
            "type": "object",
            "properties": {
                "orderId": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
