// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/cost_items/create": {
            "post": {
                "description": "Create new items data in Db.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "items"
                ],
                "summary": "Create items",
                "parameters": [
                    {
                        "description": "Create items",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserCostItems"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/cost_items/delete/{id}": {
            "delete": {
                "description": "Delete items data from Db.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "items"
                ],
                "summary": "Delete item by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Delete Items by ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/cost_items/get_all": {
            "get": {
                "description": "Get all items recorded to DB",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "items"
                ],
                "summary": "Get all items",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/cost_items/get_only_one/{id}": {
            "get": {
                "description": "Get Items By Id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "items"
                ],
                "summary": "Get Items By Id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Get Items By Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/cost_items/update/{id}": {
            "post": {
                "description": "Update items data in Db.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "items"
                ],
                "summary": "Update Items",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Update Items by ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "20": {
                        "description": "",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/daily_expense/get_by_date": {
            "get": {
                "description": "Get Expense By date",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "expense"
                ],
                "summary": "Get Expense By date",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.UserCostItems": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "item_name": {
                    "type": "string"
                }
            }
        }
    },
    "externalDocs": {
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Booker Api",
	Description:      "This is an items API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}