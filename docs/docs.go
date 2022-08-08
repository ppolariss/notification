// Package docs GENERATED BY SWAG; DO NOT EDIT
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
            "name": "Maintainer Shi Yue",
            "email": "hasbai@fduhole.com"
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
        "/": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.MessageModel"
                        }
                    }
                }
            }
        },
        "/messages": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Message"
                ],
                "summary": "List Messages of a User",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Message"
                            }
                        }
                    }
                }
            }
        },
        "/messages/clear": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Message"
                ],
                "summary": "Clear Messages of a User",
                "responses": {
                    "204": {
                        "description": ""
                    }
                }
            }
        },
        "/messages/{id}": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Message"
                ],
                "summary": "Delete a message of a user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "message id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    }
                }
            }
        },
        "/users/push-tokens": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Token"
                ],
                "summary": "List Tokens of a User",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.PushToken"
                            }
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Token"
                ],
                "summary": "Add Token of a User",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.PushToken"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Token"
                ],
                "summary": "Delete the token of a user's certain device",
                "responses": {
                    "204": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Message": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "time_created": {
                    "type": "string"
                },
                "time_updated": {
                    "type": "string"
                }
            }
        },
        "models.MessageModel": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.PushToken": {
            "type": "object",
            "properties": {
                "device_id": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "service": {
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                },
                "user_id": {
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
	Version:          "2.0.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Notification Center",
	Description:      "This is a notification microservice.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
