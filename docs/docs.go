// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "admin",
            "url": "http://subalgo.com/support",
            "email": "admin@subalgo.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/accounts": {
            "get": {
                "description": "Get all accounts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Show all accounts",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/entities.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/entities.Account"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "description": "Update an account with the provided data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Update an account",
                "parameters": [
                    {
                        "description": "Account Data",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.Account"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/entities.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "description": "Create an account with the provided data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Create a new account",
                "parameters": [
                    {
                        "description": "Account Request",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.AccountRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/entities.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete all accounts from the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Delete all accounts",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/entities.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/accounts/{id}": {
            "delete": {
                "description": "Delete an account by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Delete an account",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Account ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/entities.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "patch": {
                "description": "Update specific fields of an account by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Partially update an account",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Account ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Partial Account Request",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.AccountRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/entities.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/plans": {
            "get": {
                "description": "Retrieve all planning entries",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "plans"
                ],
                "summary": "Get all plans",
                "responses": {
                    "200": {
                        "description": "Success response with all plans information",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/entities.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/entities.PlanDetails"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/entities.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new planning entry",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "plans"
                ],
                "summary": "Create a new plan",
                "parameters": [
                    {
                        "description": "Plan information",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.PlanningRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success response with created plan information",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/entities.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/entities.PlanDetails"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/entities.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/entities.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/plans/{id}": {
            "get": {
                "description": "Retrieve a planning entry by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "plans"
                ],
                "summary": "Get a plan by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Plan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success response with plan information",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/entities.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/entities.PlanDetails"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Plan not found",
                        "schema": {
                            "$ref": "#/definitions/customerrors.CustomError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/entities.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update a planning entry by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "plans"
                ],
                "summary": "Update an existing plan",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Plan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated plan information",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.PlanningRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success response with updated plan information",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/entities.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/entities.PlanDetails"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/entities.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Plan not found",
                        "schema": {
                            "$ref": "#/definitions/entities.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/entities.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a planning entry by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "plans"
                ],
                "summary": "Delete a plan by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Plan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success response indicating plan deletion",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/entities.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/entities.PlanDetails"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Plan not found",
                        "schema": {
                            "$ref": "#/definitions/customerrors.CustomError"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/entities.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/transactions": {
            "get": {
                "description": "Retrieve transactions for a specific date",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "Get all transactions by date",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Transaction date",
                        "name": "date",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success response with list of transactions",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/entities.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/entities.Transaction"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Invalid input parameters",
                        "schema": {
                            "$ref": "#/definitions/entities.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/entities.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new transaction with the provided data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "Create a new transaction",
                "parameters": [
                    {
                        "description": "Transaction Request",
                        "name": "transaction",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.TransactionReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success response with created transaction",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/entities.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/entities.Transaction"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/entities.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/entities.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/transactions/all": {
            "delete": {
                "description": "Delete all transactions in the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "Delete all transactions",
                "responses": {
                    "200": {
                        "description": "Success response indicating all transactions deletion",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/entities.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/entities.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/transactions/{id}": {
            "delete": {
                "description": "Delete a transaction by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "Delete a transaction",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Transaction ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success response indicating deletion",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/entities.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Invalid ID format",
                        "schema": {
                            "$ref": "#/definitions/entities.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Transaction not found",
                        "schema": {
                            "$ref": "#/definitions/entities.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users/salary": {
            "get": {
                "description": "Retrieve the salary details for a specific user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user salary information",
                "responses": {
                    "200": {
                        "description": "Success response with user's salary information",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/entities.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/entities.SalaryAndResetDate"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/entities.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "customerrors.CustomError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "errorType": {
                    "type": "string"
                }
            }
        },
        "entities.Account": {
            "type": "object",
            "properties": {
                "accountId": {
                    "type": "integer"
                },
                "accountName": {
                    "type": "string"
                },
                "balance": {
                    "type": "number"
                },
                "colorIndex": {
                    "type": "integer"
                },
                "createDate": {
                    "type": "string"
                },
                "updatePlanDate": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "entities.AccountRequest": {
            "type": "object",
            "properties": {
                "accountName": {
                    "type": "string"
                },
                "balance": {
                    "type": "number"
                },
                "colorIndex": {
                    "type": "integer"
                }
            }
        },
        "entities.ErrorResponse": {
            "description": "Response containing account list",
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "errorMessage": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
                }
            }
        },
        "entities.PlanDetails": {
            "type": "object",
            "properties": {
                "accountName": {
                    "type": "string"
                },
                "amount": {
                    "type": "number"
                },
                "createDate": {
                    "type": "string"
                },
                "iconIndex": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "planId": {
                    "type": "integer"
                },
                "updateDate": {
                    "type": "string"
                },
                "usage": {
                    "type": "number"
                }
            }
        },
        "entities.PlanningRequest": {
            "type": "object",
            "properties": {
                "accountId": {
                    "type": "integer"
                },
                "amount": {
                    "type": "number"
                },
                "iconIndex": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "entities.Response": {
            "description": "Generic response format",
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "description": {
                    "type": "string"
                }
            }
        },
        "entities.SalaryAndResetDate": {
            "type": "object",
            "properties": {
                "currentUsageMonthly": {
                    "type": "number"
                },
                "resetDatePlanning": {
                    "type": "string"
                },
                "salary": {
                    "type": "number"
                }
            }
        },
        "entities.Transaction": {
            "type": "object",
            "properties": {
                "accountId": {
                    "type": "integer"
                },
                "amount": {
                    "type": "number"
                },
                "createDate": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "operation": {
                    "type": "string"
                },
                "planId": {
                    "type": "integer"
                },
                "transactionId": {
                    "type": "integer"
                },
                "updateDate": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "entities.TransactionReq": {
            "type": "object",
            "properties": {
                "accountId": {
                    "type": "integer"
                },
                "amount": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                },
                "operation": {
                    "type": "string"
                },
                "planId": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{"https", "http"},
	Title:            "User API by Fiber and Swagger",
	Description:      "API user management Server by Fiber | Doc by Swagger.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
