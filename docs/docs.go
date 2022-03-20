// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Pol Torres",
            "email": "apolinario.torresjr@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/users/{source}": {
            "post": {
                "description": "Get details of all provided github users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get details of all provided github users",
                "parameters": [
                    {
                        "description": "List of users",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/presentation.GetAccountDetailsRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Account source",
                        "name": "source",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/presentation.GetAccountDetailsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/presentation.GetAccountDetailsResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/presentation.GetAccountDetailsResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/presentation.GetAccountDetailsResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "presentation.GetAccountDetailsRequest": {
            "type": "object",
            "properties": {
                "users": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "presentation.GetAccountDetailsResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "result": {
                    "type": "boolean"
                },
                "userDetails": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/presentation.UserDetails"
                    }
                }
            }
        },
        "presentation.UserDetails": {
            "type": "object",
            "properties": {
                "company": {
                    "type": "string"
                },
                "followers": {
                    "type": "integer"
                },
                "login": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "publicRepos": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Account Detail APIs",
	Description:      "A service that provides user account details",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
