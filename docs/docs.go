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
        "/api/v1/create-portfolio": {
            "post": {
                "security": [
                    {
                        "TGWebAppToken": []
                    }
                ],
                "description": "create portfolio for user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Portfolio"
                ],
                "summary": "CreatePortfolio",
                "operationId": "create-portfolio",
                "parameters": [
                    {
                        "description": "portfolio info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreatePortfolioRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.CreatePortfolio"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.NoData"
                        }
                    }
                }
            }
        },
        "/api/v1/delete-portfolio": {
            "delete": {
                "security": [
                    {
                        "TGWebAppToken": []
                    }
                ],
                "description": "delete portfolio",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Portfolio"
                ],
                "summary": "DeletePortfolio",
                "operationId": "delete-portfolio",
                "parameters": [
                    {
                        "description": "portfolio info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.DeletePortfolioRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.NoData"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.NoData"
                        }
                    }
                }
            }
        },
        "/api/v1/portfolios/{id}": {
            "get": {
                "security": [
                    {
                        "TGWebAppToken": []
                    }
                ],
                "description": "get portfolios",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Portfolio"
                ],
                "summary": "GetPortfolios",
                "operationId": "get-portfolio",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "owner id",
                        "name": "id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.GetPortfolios"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.NoData"
                        }
                    }
                }
            }
        },
        "/api/v1/registration": {
            "post": {
                "description": "Registration new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Registration",
                "operationId": "registration",
                "parameters": [
                    {
                        "description": "user info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/utils.Body"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.Body"
                        }
                    }
                }
            }
        },
        "/api/v1/update-portfolio": {
            "patch": {
                "security": [
                    {
                        "TGWebAppToken": []
                    }
                ],
                "description": "update portfolio",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Portfolio"
                ],
                "summary": "UpdatePortfolio",
                "operationId": "update-portfolio",
                "parameters": [
                    {
                        "description": "portfolio info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdatePortfolioRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.NoData"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.NoData"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.CreatePortfolioRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "Portfolio"
                },
                "owner_id": {
                    "type": "integer",
                    "example": 518774723
                }
            }
        },
        "model.DeletePortfolioRequest": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "ownerID": {
                    "type": "integer",
                    "example": 518774723
                }
            }
        },
        "model.UpdatePortfolioRequest": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "new_name": {
                    "type": "string",
                    "example": "Portfolio"
                },
                "owner_id": {
                    "type": "integer",
                    "example": 518774723
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "firstName": {
                    "type": "string",
                    "example": "firstName"
                },
                "id": {
                    "type": "integer",
                    "example": 51
                },
                "isPremium": {
                    "type": "integer",
                    "example": 0
                },
                "lastName": {
                    "type": "string",
                    "example": "lastName"
                },
                "username": {
                    "type": "string",
                    "example": "username"
                }
            }
        },
        "schema.CreatePortfolio": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "integer",
                    "example": 1
                },
                "request": {
                    "type": "string",
                    "example": "/api/v1/endpoint"
                },
                "time": {
                    "type": "string",
                    "example": "2023-08-01T00:00:00Z"
                },
                "title": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "schema.GetPortfolios": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.Portfolio"
                    }
                },
                "request": {
                    "type": "string",
                    "example": "/api/v1/endpoint"
                },
                "time": {
                    "type": "string",
                    "example": "2023-08-01T00:00:00Z"
                },
                "title": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "schema.NoData": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string",
                    "example": "status message"
                },
                "request": {
                    "type": "string",
                    "example": "/api/v1/endpoint"
                },
                "time": {
                    "type": "string",
                    "example": "2023-08-01T00:00:00Z"
                },
                "title": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "schema.Portfolio": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2023-08-01T00:00:00Z"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "Portfolio"
                },
                "profit": {
                    "type": "number",
                    "example": 1000
                },
                "updated_at": {
                    "type": "string",
                    "example": "2023-08-01T00:00:00Z"
                }
            }
        },
        "utils.Body": {
            "type": "object",
            "properties": {
                "data": {},
                "request": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "TGWebAppToken": {
            "type": "apiKey",
            "name": "TGWebAppToken",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "TMA API",
	Description:      "API Server for Authentication",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
