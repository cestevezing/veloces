// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/api/orders": {
            "post": {
                "description": "Creates a new order with the provided details.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Create a new order",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Idempotency Key to prevent duplicate requests",
                        "name": "Idempotency-Key",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Order creation payload",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.OrderCreate"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Order created successfully",
                        "schema": {
                            "$ref": "#/definitions/model.Order"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/orders/{id}": {
            "get": {
                "description": "Retrieves an order using its unique ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Get order by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Order ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Order retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Invalid order ID",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/products": {
            "get": {
                "description": "Retrieve a list of all products.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Get All Products",
                "responses": {
                    "200": {
                        "description": "Products retrieved successfully",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Product"
                            }
                        }
                    },
                    "404": {
                        "description": "Products not found",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/products/{id}": {
            "get": {
                "description": "Retrieve a single product by its ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Get Product by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Product retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/model.Product"
                        }
                    },
                    "400": {
                        "description": "Invalid product ID",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/api/products/{id}/stock": {
            "put": {
                "description": "Update the stock quantity of a product.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Update Product Stock",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Idempotency Key to prevent duplicate requests",
                        "name": "Idempotency-Key",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Stock update payload",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.ProductStock"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Product stock updated successfully",
                        "schema": {
                            "$ref": "#/definitions/model.Product"
                        }
                    },
                    "400": {
                        "description": "Invalid request body or product ID",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Order": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "customerName": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.OrderItem"
                    }
                },
                "totalAmount": {
                    "type": "number"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "model.OrderItem": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "orderID": {
                    "type": "integer"
                },
                "productID": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                },
                "subtotal": {
                    "type": "number"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "model.Product": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "stock": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "requests.OrderCreate": {
            "type": "object",
            "required": [
                "customer_name",
                "items"
            ],
            "properties": {
                "customer_name": {
                    "type": "string"
                },
                "items": {
                    "type": "array",
                    "minItems": 1,
                    "items": {
                        "$ref": "#/definitions/requests.OrderItemCreate"
                    }
                }
            }
        },
        "requests.OrderItemCreate": {
            "type": "object",
            "required": [
                "product_id",
                "quantity"
            ],
            "properties": {
                "product_id": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "requests.ProductStock": {
            "type": "object",
            "required": [
                "new_stock"
            ],
            "properties": {
                "new_stock": {
                    "type": "integer"
                }
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
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
